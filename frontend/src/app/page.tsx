import LinkComponent from "./features/components/elements/link/Link";
import { UrlArray } from "./features/constants/Url";
import { UrlMsgArray } from "./features/constants/Urlmsg";
import { CiMemoPad } from "react-icons/ci";

export default async function Home() {
  return (
    <div className="flex h-screen flex-col items-center justify-center">
        <div className="mx-4 mb-8 flex items-center justify-center text-5xl">
          Welcome Todo App!
          <CiMemoPad className="text-6xl"/>
        </div>
        <div className="flex justify-center gap-4">
          <LinkComponent href={UrlArray[0]}>{UrlMsgArray[0]}</LinkComponent>
          <LinkComponent href={UrlArray[1]}>{UrlMsgArray[1]}</LinkComponent>
        </div>
    </div>
  )
}
