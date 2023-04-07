import { React, useRef, useContext } from 'react'
import './about.css'
// import ME from '../../assets/me.jpg'
import { FaDocker, FaReact, FaAws } from 'react-icons/fa'
import { DiNodejs } from 'react-icons/di'
import {
    SiNginx, SiPython, SiDjango, SiFastapi,
    SiFlask, SiGithub, SiKubernetes, SiVim, SiScala,
    SiJava, SiCsharp, SiRedis, SiPostgresql, SiJupyter, SiNumpy, SiPandas, SiTensorflow, SiKeras, SiApachespark, SiPlotly,
} from 'react-icons/si'


import UserContext from '../../UserContext';

const AboutMeItem = ({ icon, text }) => {
    const { user } = useContext(UserContext);
    const alreadyHoveredRef = useRef(false)
    const handleHover = () => {
        if (alreadyHoveredRef.current) return
        alreadyHoveredRef.current = true

        fetch('http://localhost:8080/log-grpc', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                text, // or any other data you want to send in the post request
                user
            })
        }).then(response => {
            console.log(response)
        }).catch(error => {
            console.log(error)
        })
    }
    const handleMouseLeave = () => {
        alreadyHoveredRef.current = false;
    };

    return (
        <div className='about__me-img' onMouseOver={handleHover} onMouseLeave={handleMouseLeave}>
            {icon}
            {text}
        </div>
    )
}

const About = () => {

    return (
        <section id='about'>
            <h2 >I've been doing projects using ...</h2>

            <div className='container about__container'>
                <div className='about__me'>
                    <h2>Machine Learning
                    </h2>
                    <br></br>
                    <div className='about__mes'>
                        <AboutMeItem icon={<SiPython className="icon" />} text="Python" />
                        <AboutMeItem icon={<SiJupyter className="icon" />} text="Jupyter" />
                        <AboutMeItem icon={<SiKeras className="icon" />} text="Keras" />
                        <AboutMeItem icon={<SiTensorflow className="icon" />} text="Tensorflow" />
                        <AboutMeItem icon={<SiPandas className="icon" />} text="Pandas" />
                        <AboutMeItem icon={<SiNumpy className="icon" />} text="Numpy" />
                        <AboutMeItem icon={<h2 className="icon">ResNet</h2>} />
                        <AboutMeItem icon={<h2 className="icon">NASNet</h2>} />
                        <AboutMeItem icon={<h2 className="icon">BERT</h2>} />
                        <AboutMeItem icon={<SiApachespark className="icon" />} text="Apache Spark" />
                        <AboutMeItem icon={<SiScala className="icon" />} text="Scala" />
                        <AboutMeItem icon={<SiPlotly className="icon" />} text="Plotly" />
                    </div>

                </div>
                <div className='about__me'>
                    <h2 >Software Engineering</h2>
                    <br></br>
                    <div className='about__mes'>
                        <AboutMeItem icon={<SiPython className="icon" />} text="Python" />
                        <AboutMeItem icon={<DiNodejs className="icon" />} text="Nodejs" />
                        <AboutMeItem icon={<SiDjango className="icon" />} text="Django" />
                        <AboutMeItem icon={<SiFastapi className="icon" />} text="Fastapi" />
                        <AboutMeItem icon={<SiFlask className="icon" />} text="Flask" />
                        <AboutMeItem icon={<FaDocker className="icon" />} text="Docker" />
                        <AboutMeItem icon={<SiKubernetes className="icon" />} text="Kubernetes" />
                        <AboutMeItem icon={<FaAws className='icon' />} text="AWS" />
                        <AboutMeItem icon={<SiGithub className='icon' />} text="GitHub" />
                        <AboutMeItem icon={<SiVim className="icon" />} text="Vim" />
                        <AboutMeItem icon={<SiScala className="icon" />} text="Scala" />
                        <AboutMeItem icon={<SiJava className="icon" />} text="Java" />
                        <AboutMeItem icon={<SiCsharp className="icon" />} text="C#" />
                        <AboutMeItem icon={<SiRedis className="icon" />} text="Redis" />
                        <AboutMeItem icon={<SiPostgresql className="icon" />} text="PostgreSQL" />
                        <AboutMeItem icon={<SiNginx className="icon" />} text="Nginx" />
                        <AboutMeItem icon={<FaReact className="icon" />} text="React" />

                    </div>
                </div>


                <div className='about__content'>
                    <div className='about__cards'>
                        <article className='about__card' >
                            <h5 className='about__card__head'>About experience in Machine learing:</h5>
                            <small className='text'>
                                {"\n"}
                                I am currently a MSDS student in USF, focusing on production-level Machine Learning practices.{"\n"}
                                I've been build both NASnet and ResNet with Tensorflow as a final project in school(undergrad).{"\n"}
                                I applied BERT model for NLP Recommendation system for start-up company that I co-founded.{"\n"}

                            </small>

                        </article>
                    </div>

                    <div className='about__cards'>
                        <article className='about__card' >
                            <h5 className='about__card__head'>About software engineering experience:</h5>
                            <small className='text'>{"\n"} As a self-taught software engineer with over one year experience.
                                I'm really comfortable with all sorts of learning curve(I'm a big fan of Vim by the way).{"\n"}
                                Though I have experience with Java, C#, and Scala, my area of expertise is mainly Python and Javascript.{"\n"}
                            </small>
                        </article>
                    </div>
                </div>

            </div>
        </section >
    )
}

export default About