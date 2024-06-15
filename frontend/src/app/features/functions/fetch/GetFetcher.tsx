export default async function GetFetcher(endpoint: string){
  return await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}${endpoint}`, {
    method: 'GET',
    //body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json',
    },
  })
}