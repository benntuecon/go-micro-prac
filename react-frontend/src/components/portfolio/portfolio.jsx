
import React from 'react'
import './portfolio.css'
import DASH from '../../assets/dash.png'
import tree from '../../assets/tree.png'
import santorini from '../../assets/santorini.jpeg'



const Portfolio = () => {
    // golanf project is at localhost:8080 make a post request to localhost:8080/
    const [goProjects, setGoProject] = React.useState('')

    const handleGoProject = () => {
        let body = {
            method: 'POST',
        }
        fetch('http://localhost:8080/', body)
            .then(res => res.json())
            .then(data => setGoProject(data.msg))
            .catch(err => console.log(err))
    }

    const handleGoProjectAuth = () => {
        let body = {
            action: "auth",
            auth: {
                email: "admin@example.com",
                password: "verysecret",
            }
        }
        const headers = new Headers()
        headers.append('Content-Type', 'application/json')
        body = JSON.stringify(body)

        let payload = {
            method: 'POST',
            body,
            headers
        }

        fetch('http://localhost:8080/handle', payload)
            .then(res => res.json())
            .then(data => setGoProject(data.msg))
            .catch(err => console.log(err))
    }

    return (
        <section id='portfolio'>
            <h2>Selected Projects</h2>
            <div className='container portfolio__container'>
                <article className='portfolio__item'>
                    <div className='portfolio__item-image'>
                        <img src={DASH} alt='' />
                    </div>
                    <h3>Python Fastapi Dash webapp</h3>
                    <a href="https://github.com/benntuecon/dash-website" className='btn' target='_blank' rel='noreferrer'>GitHub</a>

                </article>
                <article className='portfolio__item'>
                    <div className='portfolio__item-image'>
                        <img src={tree} alt='' />
                    </div>
                    <h3>Bagging neuron network </h3>
                    <a href="https://www.kaggle.com/code/bennbenbenn/random-vnn-forest-credit-card-fraud/notebook" className='btn' target='_blank'>Kaggle</a>

                </article>
                <article className='portfolio__item'>
                    <div className='portfolio__item-image'>
                        <img src={santorini} alt='' />
                    </div>
                    <h3>Santorini broadgame Java clone(working)</h3>
                    <a href="https://github.com/benntuecon/JAVA_santorini" className='btn' target='_blank'>GitHub</a>

                </article>
                <article className='portfolio__item'>
                    <div className='portfolio__item-image'>
                        <img src={santorini} alt='' />
                    </div>
                    <h3>Golang micro service {goProjects}</h3>
                    <a href="https://github.com/benntuecon/go-micro-prac" className='btn' target='_blank'>GitHub</a>

                    <button onClick={handleGoProject}>broker</button>
                    <button onClick={handleGoProjectAuth}>Auth</button>

                </article>

            </div>
        </section>
    )
}

export default Portfolio