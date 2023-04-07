import React from 'react'
import About from './components/about/About'
import Contact from './components/contact/Contact'
import Header from './components/header/Header'
import Nav from './components/nav/nav'
import Portfolio from './components/portfolio/portfolio'
import UserProvider from './UserProvider'



const App = () => {
    return (
        <>
            <UserProvider>
                <Header />
                <Nav />
                <About />
                <Portfolio />
                <Contact />
            </UserProvider>
        </>
    )
}

export default App