import microservicesImage from '../../assets/images/icon2.png';

export function HeaderSection() {
    const handleClick = (event: React.MouseEvent<HTMLAnchorElement>) => {
        event.preventDefault();
        const aboutSection = document.getElementById('about');
        aboutSection?.scrollIntoView({ behavior: 'smooth' });
    };

    return (
        <div className="bg-white flex">
            <div className="w-1/2">
                <div className="relative isolate px-6 pt-14 lg:px-8">
                    <div className="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
                        <div className="hidden sm:mb-8 sm:flex sm:justify-center">
                            <div className="relative rounded-full px-3 py-1 text-sm leading-6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20">
                                Tell me more about microservices! {' '}
                                <a className="font-semibold text-selectiveYellow" onClick={handleClick}>
                                    <span className="absolute inset-0" aria-hidden="true" />
                                    Read more <span aria-hidden="true">&rarr;</span>
                                </a>
                            </div>
                        </div>
                        <div className="text-center">
                            <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
                                Serverless Orchestrator
                            </h1>
                            <p className="mt-6 text-lg leading-8 text-gray-600">
                                A microservice platform enhancing software efficiency
                                through scalable, independent service deployability for streamlined development.
                            </p>
                            <div className="mt-10 flex items-center justify-center gap-x-6">
                                <a
                                    href="/GetStarted"
                                    className="rounded-md bg-darkPink px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-amaranthPink focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                >
                                    Get started
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className="w-1/2">
                <img src={microservicesImage} alt="Microservice Platform Image" />
            </div>
        </div>
    );
}
