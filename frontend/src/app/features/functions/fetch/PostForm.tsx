import type { FormValues } from '../../types/data'

export default async function PostForm(data: FormValues, endpoint: string) {
  return await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}${endpoint}`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json',
    },
  })
}
