"use client"

import type { SubmitHandler } from "react-hook-form"
import type { FormValues } from "@/app/features/types/data"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { PostForm } from "@/app/features/functions"
import { useRouter } from "next/navigation"
import { z } from "zod"
import * as Url from "@/app/features/constants";
import LinkComponent from "@/app/features/components/elements/link/Link"
import { Button } from "@/app/features/components/elements"

const signUpSchema = z
    .object({
        name: z
            .string()
            .min(1, "名前は1文字以上入力してください"),
        email: z
            .string()
            .email("正しいメールアドレスを入力してください"),
        password: z
            .string()
            .min(8, "パスワードは8文字以上入力してください")
            .regex(/^[a-zA-Z0-9]+$/, "パスワードは英数字のみで入力してください"),
});

export default function SignUpPage() {
    const router = useRouter()
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<FormValues>({
        resolver: zodResolver(signUpSchema)
    })

    const onsubmit: SubmitHandler<FormValues> = async (data) => {
        try{
            const res = await PostForm(data, "signup")
            console.log(res)
            router.push("/")
        }catch (err) {
            console.log(err)
        }
    }

    return (
        <div className="flex h-screen flex-col items-center justify-center">
            <form onSubmit={handleSubmit(onsubmit)} className="w-96 rounded-lg bg-white p-8 shadow-md">
                <h1 className="mb-4 text-2xl font-medium text-gray-700">新規登録</h1>
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-600">Name</label>
                    <input type="text" {...register("name")} className="mt-1 w-full rounded-md border-2 p-2"/>
                </div>
                {errors.name && (<span className="text-sm text-red-600">{errors.name.message}</span>)}
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-600">Email</label>
                    <input type="text" {...register("email")} className="mt-1 w-full rounded-md border-2 p-2"/>
                </div>
                {errors.email && (<span className="text-sm text-red-600">{errors.email.message}</span>)}
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-600">PassWord</label>
                    <input type="password" {...register("password")} className="mt-1 w-full rounded-md border-2 p-2"/>
                </div>
                {errors.password && (<span className="text-sm text-red-600">{errors.password.message}</span>)}

                <div className="flex justify-end">
                    <Button type="submit">新規登録</Button>
                </div>
                <LinkComponent href={Url.UrlArray[3]}>{Url.UrlMsgArray[1]}</LinkComponent>
            </form>
        </div>
    )
}