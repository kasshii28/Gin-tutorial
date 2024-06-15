import { Button } from "./features/components/elements";
import { UrlArray } from "./features/constants/Url";
import { UrlMsgArray } from "./features/constants/Urlmsg";

export default async function Home() {
  return (
    <div>
      <Button url={UrlArray[0]} msg={UrlMsgArray[0]}/>
      <Button url={UrlArray[1]} msg={UrlMsgArray[1]}/>
    </div>
  )
}
