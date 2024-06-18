import type { Metadata } from 'next'
import MainPage from "@/components/pages/MainPage";


// export const metadata: Metadata = {
//   title: "Ural Radio Enjeenering journal",
//   description: "Рецензируемый международный научно-технический журнал с открытым доступом, посвященный последним достижениям радиоэлектроники и связи.",
// };

export default function Home() {
  return (
    <>
      <MainPage/>
    </>
  )
}
