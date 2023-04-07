import React from 'react'
import About from './components/about/About'
import Contact from './components/contact/Contact'
import Header from './components/header/Header'
import Nav from './components/nav/nav'
import Portfolio from './components/portfolio/portfolio'


const App = () => {
    return (
        <>
            <Header />
            <Nav />
            <About />
            <Portfolio />
            <Contact />
        </>
    )
}

export default App