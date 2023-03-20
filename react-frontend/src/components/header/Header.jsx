import React from 'react'
import './header.css'
import CTA from './CTA'
import ME from '../../assets/me.jpg'
import HeaderSocials from './HeaderSocials'

const Header = () => {
    return (
        <header>
            <div className="container header__container">
                <h5>
                    Hello, this is
                </h5>
                <h1>
                    👋🏼 Ben Chen
                </h1>
                <h5 >
                    A Data Scientist & Software Engineer, currently a Master of Science in Data Science student at <a href="https://www.usfca.edu/arts-sciences/programs/graduate/data-science/" target="_blank" rel="noreferrer" >USF</a>
                </h5>
                <h5 className='text-light'>

                </h5>
                <CTA />
                <HeaderSocials />
                <div className='me'>
                    <img src={ME} alt='me' />
                </div>

                <a href='#contact' className='scroll__down'> Scroll Down</a>
            </div>

        </header >
    )
}

export default Header