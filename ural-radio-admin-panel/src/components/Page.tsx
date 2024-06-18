import React from 'react';

const Page = ({ children }: {children: React.ReactNode}) => {
  return (
    <>
      <main className='main'>
        <div className="main__container">
          <div className="main__title">
            <div className="main__title-line">URAL</div>
            <div className="main__title-line">RADIO ENGINEERING</div>
            <div className="main__title-line">JOURNAL</div>
          </div>
            
          {children}
        </div>
      </main>
      <div className="circle1"></div>
      <div className="circle2"></div>
      <div className="circle3"></div>
    </>
  );
};

export default Page;