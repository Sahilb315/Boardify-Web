"use client";
import * as React from "react";
import { FieldValue, FieldValues, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import IconOpenEye from "@/components/Icons/IconOpenEye";
import { useState } from "react";
import IconCloseEye from "@/components/Icons/IconCloseEye";
import Link from "next/link";
import axios from "axios";
import { BASE_URL } from "../../../../constants";
import { showToast } from "@/components/ShowToast";
import { useRouter } from "next/navigation";

const schema = z.object({
  fullName: z.string().min(3, "Name must be at least 3 characters"),
  email: z.string().email({ message: "Invalid email address" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters" }),
});

const SignUp = () => {
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(schema),
  });

  const router = useRouter();

  const onSubmit = async (data: FieldValues) => {
    setIsLoading(true);
    try {
      let req = await axios.post(`${BASE_URL}/sign_up`, {
        full_name: data.fullName,
        email: data.email,
        password: data.password,
      });
      if (req.status === 201) {
        setIsLoading(false);
        showToast("Sign Up Successful", "success");
        router.push("/login");
      }
    } catch (error) {
      let errorMessage = "An unexpected error occurred. Please try again.";
      if (axios.isAxiosError(error)) {
        if (error.response) {
          switch (error.response.status) {
            case 400:
              errorMessage = "Invalid request. Please check your input.";
              break;
            case 401:
              errorMessage =
                "You are not authorized to view this data. Please log in.";
              break;
            case 404:
              errorMessage = "The requested data could not be found.";
              break;
            case 750:
              errorMessage =
                "We're experiencing server issues. Please try again later.";
              break;
          }
        } else if (error.request) {
          errorMessage =
            "Unable to connect to the server. Please check your internet connection.";
        }
      }
      setIsLoading(false);
      showToast(errorMessage, "error");
    }
  };

  return (
    <div className="relative flex items-center justify-center h-screen bg-gradient-to-t from-purple to-white">
      {isLoading && (
        <>
          <div className="fixed inset-0 flex z-50 items-center justify-center bg-gray-100 bg-opacity-80">
            <span className="inline-block h-12 w-12 animate-spin rounded-full border-4 border-[#4534AC]/70 border-t-transparent align-middle"></span>
          </div>
        </>
      )}
      <Card className="w-[450px]">
        <CardHeader>
          <CardTitle className="text-center mb-4">
            Welcome to <span className="text-[#4534AC]">WorkFlo!</span>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="fullName">Full Name</Label>
                <Input
                  disabled={isLoading}
                  id="fullName"
                  placeholder="Full Name"
                  {...register("fullName")}
                />
                {errors.fullName && (
                  <p className="text-red-500 mt-1 text-sm">
                    {errors.fullName.message?.toString()}
                  </p>
                )}
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="email">Email</Label>
                <Input
                  disabled={isLoading}
                  id="email"
                  placeholder="Your Email"
                  {...register("email")}
                />
                {errors.email && (
                  <p className="text-red-500 mt-1 text-sm">
                    {errors.email.message?.toString()}
                  </p>
                )}
              </div>
              <div className="">
                <Label
                  htmlFor="password"
                  className={`mb-2 block ${isLoading ? "opacity-50" : ""}`}
                >
                  Password
                </Label>
                <div className="relative">
                  <Input
                    disabled={isLoading}
                    id="password"
                    type={`${isPasswordVisible ? "text" : "password"}`}
                    placeholder="Password"
                    className={`pr-10 ${isLoading ? "opacity-50" : ""}`}
                    {...register("password")}
                  />
                  <div
                    onClick={() => setIsPasswordVisible(!isPasswordVisible)}
                    className={`absolute cursor-pointer inset-y-0 right-0 flex ${
                      isLoading ? "opacity-50" : ""
                    } items-center pr-3`}
                  >
                    {isPasswordVisible ? (
                      <IconOpenEye className="h-5 w-5 text-gray-400" />
                    ) : (
                      <IconCloseEye className="h-5 w-5 text-gray-400" />
                    )}
                  </div>
                </div>
                {errors.password && (
                  <p className="text-red-500 mt-1 text-sm">
                    {errors.password.message?.toString()}
                  </p>
                )}
              </div>
            </div>
            <CardFooter className="flex flex-col justify-between mt-4">
              <Button
                type="submit"
                className="w-full bg-gradient-to-b from-[#9C93D4] to-[#4B36CC]/85 text-white hover:text-white"
                variant="outline"
              >
                Sign Up
              </Button>
              <span className="text-center mt-4">
                Already have an account?{" "}
                <Link href={"/login"}>
                  <span className="text-[#0054A1] cursor-pointer">Log in</span>
                </Link>
              </span>
            </CardFooter>
          </form>
        </CardContent>
      </Card>
    </div>
  );
};

export default SignUp;
