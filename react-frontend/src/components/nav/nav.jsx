import React from 'react'
import './nav.css'
import { AiOutlineHome } from 'react-icons/ai'
import { FaUserTie } from 'react-icons/fa'
import { AiFillTool } from 'react-icons/ai'
import { AiFillMessage } from 'react-icons/ai'
import { useState } from 'react'



const Nav = () => {
    const [activeNav, setActiveNav] = useState('#')
    return (

        <nav>
            <a href="#" onClick={() => { setActiveNav('#') }} className={activeNav === '#' ? 'active' : ''}><AiOutlineHome /></a>
            <a href="#about" onClick={() => setActiveNav('#about')} className={activeNav === '#about' ? 'active' : ''}><FaUserTie /></a>
            <a href="#portfolio" onClick={() => setActiveNav('#portfolio')} className={activeNav === '#portfolio' ? 'active' : ''}><AiFillTool /></a>
            <a href="#contact" onClick={() => setActiveNav('#contact')} className={activeNav === '#contact' ? 'active' : ''}> <AiFillMessage /></a>
        </nav >
    )
}

export default Nav