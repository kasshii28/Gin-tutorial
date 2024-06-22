import Link from "next/link"
import type { LinkProps } from "@/app/features/types/data"

export default function LinkComponent({ ...props }: LinkProps){
  return (
    <div className="mt-4">
        <Link href={props.href}>
            <span className="text-sm font-bold text-blue-500 hover:text-blue-800">
                {props.children}
            </span>
        </Link>
    </div>
  )
}