import { type Data } from "./types/data"

export default async function Home() {
  const res = await fetch("http://localhost:8000/", {
    cache: "no-store",
  })
  const data: Data = await res.json()
  return (
    <div>
      {data.message}
    </div>
  )
}
