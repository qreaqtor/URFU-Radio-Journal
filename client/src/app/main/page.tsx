import { Metadata } from "next";
import Head from "next/head";

export const metadata: Metadata = {
  title: "Ural Radio Enjeenering journal",
  description: "Рецензируемый международный научно-технический журнал с открытым доступом, посвященный последним достижениям радиоэлектроники и связи.",
};

const MainPage = () => {
  return (
    <>
      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>О журнале</p></div>
        <div className="paragraph__text">
          <p>Рецензируемый международный научно-технический журнал с открытым доступом,посвященный последним достижениям радиоэлектроники и связи.</p>
          <p>Редакция журнала отдаёт предпочтение практическим исследованиям, связанным с освоением критических промышленных технологий, реализация которых обеспечит создание высокоэффективнойконкурентоспособной радиоэлектронной продукции.</p>
          <p>Журнал зарегистрирован Федеральной службой по надзору в сфересвязи, информационных технологий и массовых коммуникаций. Свидетельство о регистрации ПИ № ФС77-69790 от 18.05.2017</p>
        </div>
      </div>

      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>Журнал ориентирован на научные специальности:</p></div>
        <div className="paragraph__text">
          <p><span className='paragraph__number'>2.2.2.</span> Электронная компонентная база микро- и наноэлектроники, квантовых устройств (технические науки)</p>
          <p><span className='paragraph__number'>2.2.8.</span> Методы и приборы контроля и диагностики материалов, изделий, веществ и природной среды (технические науки)</p>
          <p><span className='paragraph__number'>2.2.13.</span> Радиотехника, в том числе системы и устройства телевидения (технические науки)</p>
          <p><span className='paragraph__number'>2.2.14.</span> Антенны, СВЧ-устройства и их технологии</p>
          <p><span className='paragraph__number'>2.2.15.</span> Системы, сети и устройства телекоммуникаций (технические науки)</p>
          <p><span className='paragraph__number'>2.2.16.</span>  Радиолокация и радионавигация (технические науки)</p>
        </div>
        <div className="paragraph__text">
          <p>Включен в Объединенный каталог «Пресса России». Индекс 33049 Журнал входит в Перечень рецензируемых научных изданий (с 15.04.2021), рекомендованных ВАК для публикации основных научных результатов диссертаций на соискание ученой степени кандидата наук, на соискание ученой степени доктора наук.</p>
          <p>Полнотекстовая версия журнала находится в режиме свободного доступа: на сайте журнала, в электронном научном архиве УрФУ и на платформе Научной электронной библиотеки (РИНЦ).</p>
          <p>Выходит с 2017 года</p>
        </div>
      </div>
    </>
  );
};


export default MainPage;