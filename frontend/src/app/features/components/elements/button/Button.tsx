import Link from "next/link"

export default function Button({url, msg}: {url: string, msg: string}) {
  return (
    <button className="ml-1 text-sm font-bold text-blue-500 hover:text-blue-700">
        <Link href={url}>
            {msg}
        </Link>
    </button>
  )
}