export function FeaturesSection() {
    // Assume each feature has a title and description
    const features = [
        { title: 'Feature 1', description: 'Description 1' },
        { title: 'Feature 2', description: 'Description 2' },
        { title: 'Feature 1', description: 'Description 1' },
        { title: 'Feature 2', description: 'Description 2' },
        { title: 'Feature 1', description: 'Description 1' },
        { title: 'Feature 2', description: 'Description 2' },
        // ... add all features
    ];

    return (
        <div className="bg-white py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-screen-xl mx-auto">
                <div className="text-center">
                    <h2 className="text-3xl leading-9 font-extrabold text-white">
                        Features
                    </h2>
                    <p className="mt-3 max-w-2xl mx-auto text-xl leading-7 text-gray-500 sm:mt-4">
                        Here are some of the features that Serverless Orchestrator offers.
                    </p>
                </div>

                <div className="mt-10">
                    <ul className="md:grid md:grid-cols-2 lg:grid-cols-3 md:gap-x-8 md:gap-y-10">
                        {features.map((feature) => (
                            <li key={feature.title} className="mt-10 md:mt-0 bg-black border-8 border-white rounded-lg p-4 shadow">
                                <div className="flex">
                                    <div className="flex-shrink-0">
                                        <div className="h-12 w-12 bg-gray-200"></div> {/* Placeholder box */}
                                    </div>
                                    <div className="ml-4">
                                        <h5 className="text-lg leading-6 font-medium text-gray-900">{feature.title}</h5>
                                        <p className="mt-2 text-base leading-6 text-gray-500">{feature.description}</p>
                                    </div>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>

                <div className="mt-8 text-center">
                    <a href="#documentation" className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-black hover:bg-opacity-90 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out">
                        Documentation
                    </a>
                    <a href="#capstone-poster" className="ml-4 inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-black bg-white hover:bg-gray-100 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out">
                        Capstone Poster
                    </a>
                </div>
            </div>
        </div>
    );
}