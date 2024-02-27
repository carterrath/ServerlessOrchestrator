import githubIcon from '../../assets/svg/github.svg';
import boxIcon from '../../assets/svg/box.svg';
import codeIcon from '../../assets/svg/code.svg';

export function AboutSection() {
    // Assume each feature has a title and description
    const features = [
        { title: 'Simplifying Full-Stack Development', description: 'Tackle multiple domains without the complexity of managing integrated components.' , icon: <img src={githubIcon} alt="Github Icon" width="50" height="50" style={{ fill: 'white' }} /> },
        { title: 'Monolithic to Microservices', description: 'Transition from monolithic architectures to scalable, serverless models.', icon: <img src={boxIcon} alt="Github Icon" width="50" height="50" style={{ fill: 'white' }} /> },
        { title: 'Serverless Orchestration', description: 'Automate serverless functions and microservices with a container orchestration platform.' , icon: <img src={codeIcon} alt="Github Icon" width="50" height="50" style={{ fill: 'white' }} /> },
        { title: 'Developer-Friendly Environment', description: 'Supports various programming languages.' , icon: <img src={boxIcon} alt="Github Icon" width="50" height="50" style={{ fill: 'white' }} /> },
        { title: 'GitHub Integration', description: 'Easy cloning from GitHub repository for seamless collaboration.' , icon: <img src={boxIcon} alt="Github Icon" width="50" height="50" style={{ fill: 'white' }} />}, 
        { title: 'Focus on Code', description: 'Empowers developers to concentrate on business logic rather than infrastructure management.' , icon: <img src={boxIcon} alt="Github Icon" width="50" height="50" style={{ fill: 'white' }} /> },
       
    ];

    return (
      <div className="bg-white py-12 px-4 sm:px-6 lg:px-8 border-b border-slate-200">
        <div className="bg-white py-12 px-4 sm:px-6 lg:px-8">
          
            <div className="max-w-screen-xl mx-auto">
                <div className="text-center">
                    <h2 className="text-3xl font-bold leading-9 tracking-tighter sm:text-3xl md:text-6xl">
                        Features
                    </h2>
                    <p className="mt-3 max-w-2xl mx-auto text-xl leading-7 text-gray-500 sm:mt-4">
                        Here are some of the features that Serverless Orchestrator offers.
                    </p>
                </div>

                <div className="mt-10">
                    <ul className="md:grid md:grid-cols-2 lg:grid-cols-3 md:gap-x-8 md:gap-y-10">
                        {features.map((feature) => (
                            <li key={feature.title} className="mt-10 md:mt-0 bg-caribbeanCurrent border-8 border-white rounded-lg p-4 shadow">

                                <div className="flex">
                                    <div className="flex-shrink-0">
                                        {feature.icon} 
                                    </div>
                                    <div className="ml-4">
                                        <h5 className="text-lg leading-6 font-medium text-white">{feature.title}</h5>
                                        <p className="mt-2 text-base leading-6 text-white">{feature.description}</p>
                                    </div>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>

                <div className="mt-8 text-center">
                    <a href="#documentation" className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-caribbeanCurrent hover:bg-lightTeal focus:outline-none focus:shadow-outline transition duration-150 ease-in-out">
                        Documentation
                    </a>
                    <a href="#documentation" className="ml-4 inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-caribbeanCurrent hover:bg-lightTeal focus:outline-none focus:shadow-outline transition duration-150 ease-in-out">
                        Capstone Poster
                    </a>
                </div>
            </div>
        </div>
        </div>
    );
}