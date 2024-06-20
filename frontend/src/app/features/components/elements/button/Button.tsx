import type { ButtonProps } from "@/app/features/types/data";

export default function Button({ ...props }: ButtonProps) {
  return (
    <button 
      type={props.type}
      className="rounded bg-blue-500 px-4 py-2 text-sm font-bold text-white hover:bg-blue-700">
        {props.children}
    </button>
  )
}