export function AboutSection() { 
    const divs = [
        { text: 'Text 1', icon: <MyComponent /> },
        { text: 'Text 2', icon: <MyComponent /> },
        { text: 'Text 3', icon: <MyComponent /> },
      ];
    return(
        <div className="flex justify-center">
        <div className="bg-white p-6 w-full max-w-7xl text-center">
          <div className="py-24 sm:py-32">
            {/* Start of grid container */}
            <div className="mx-auto grid gap-x-8 gap-y-20 px-6 lg:px-8">
              {/* About section */}
              <div className="mx-auto">
                <h2 className="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">About</h2>
                <p className="mt-6 text-lg leading-8 text-gray-600">
                  Submit your microservice to our platform and we will handle the rest.
                </p>
              </div>
              {/* Start of list */}
              <ul role="list" className="grid gap-x-8 gap-y-12 sm:grid-cols-2 lg:grid-cols-3 xl:col-span-3 mx-auto">
                {divs.map((div, index) => (
                  <li key={index} className="flex items-center gap-x-6 rounded-lg p-4">
                    {div.icon}
                    <p>{div.text}</p>
                  </li>
                ))}
              </ul>
              {/* End of list */}
              <div className="mt-10 flex items-center justify-center gap-x-6">
                <a
                  href="#"
                  className="w-50 rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Documentation
                </a>
                <a
                  href="#"
                  className="w-50 rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  GitHub
                </a>
              </div>
            </div>
            {/* End of grid container */}
          </div>
        </div>
      </div>
    );

}

function MyComponent() {
    return (
      <div>
        Hello, this is my component!
      </div>
    );
  }