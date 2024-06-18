import React from 'react';

const Footer = () => {
  const data = new Date();
  return (
    <footer className='footer'>
      <div className="footer__text">{data.getFullYear()} г. Издательство Уральского университета Россия, 620083, Екатеринбург, ул. Тургенева, 4</div>
    </footer>
  );
};

export default Footer;