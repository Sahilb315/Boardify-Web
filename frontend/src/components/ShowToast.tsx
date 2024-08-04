import { toast } from "sonner";

export const showToast = (
  message: string,
  type: "success" | "error" | "normal"
) => {
  switch (type) {
    case "success":
      toast.success(message, {
        description: "   ",
        duration: 2000,
      });
      break;
    case "error":
      toast.error(message, {
        description: "   ",
        duration: 2000,
      });
      break;
    case "normal":
      toast(message, {
        description: "   ",
        duration: 2000,
      });
      break;
    default:
      toast(message, {
        description: "   ",
        duration: 2000,
      });
  }
};
