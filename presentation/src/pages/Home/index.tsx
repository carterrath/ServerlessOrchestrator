
import React, { useState } from 'react';
import microservicesImage from '../../assets/images/icon2.png';
import photoMae from '../../assets/images/photoMae.png';
import photoRuth from '../../assets/images/photoRuth.png';
import photoCarter from '../../assets/images/photoCarter.png';
import photoJackie from '../../assets/images/photoJackie.png';


export function Home() {
  const [isTop, setIsTop] = useState(true);

  const handleScroll = () => {
    const position = window.pageYOffset;
    setIsTop(position < 50);
  };

  const divs = [
    { text: 'Text 1', icon: <MyComponent /> },
    { text: 'Text 2', icon: <MyComponent /> },
    { text: 'Text 3', icon: <MyComponent /> },
  ];

  const people = [
    {
      name: 'Mae Pereyra',
      role: 'Software Barbie',
      imageUrl: photoMae,
      linkedinUrl: 'url_to_linkedin_profile',
    },
    {
      name: 'Ruth Jimenez',
      role: 'Software Barbie',
      imageUrl: photoRuth,
      linkedinUrl: 'url_to_linkedin_profile',
    },
    {
      name: 'Carter Rath',
      role: 'Software Barbie',
      imageUrl: photoCarter,
      linkedinUrl: 'url_to_linkedin_profile',
    },
    {
      name: 'Jaclyn Walsh',
      role: 'Software Barbie',
      imageUrl: photoJackie,
      linkedinUrl: 'https://www.linkedin.com/in/jaclynewalsh/',
    },
    
  ];

  React.useEffect(() => {
    window.addEventListener('scroll', handleScroll, { passive: true });

    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []);

  return (
    <div className="flex flex-col">
      <div className="bg-white flex">
        <div className="w-1/2 border-b border-slate-200">
        <div className="relative isolate px-6 pt-14 lg:px-8">
            <div className="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
              aria-hidden="true">
              <div className="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
                style={{
                  clipPath:
                    'polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)',
                }}
              />
            </div>
          <div className="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
            <div className="hidden sm:mb-8 sm:flex sm:justify-center">
              <div className="relative rounded-full px-3 py-1 text-sm leading-6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20">
                Announcing our next round of funding.{' '}
                <a href="#" className="font-semibold text-indigo-600">
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
                  href="#"
                  className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Get started
                </a>
              
              </div>
            </div>
          </div>
          <div
            className="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
            aria-hidden="true"
          >
            <div
              className="relative left-[calc(50%+3rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
              style={{
                clipPath:
                  'polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)',
              }}
            />
          </div>
        </div>
        </div>
        <div className="w-1/2 border-b border-slate-200">
          <img src={microservicesImage} alt="Microservice Platform Image" />
        </div>
      </div>
      
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

      
     

      <div className="flex justify-center">
        <div className="bg-white p-6 w-full max-w-7xl">
          <div className="py-24 sm:py-32">
            <div className="mx-auto grid gap-x-8 gap-y-20 px-6 lg:px-8 xl:grid-cols-3">
              <div className="max-w-2xl">
                <h2 className="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">Meet our team</h2>
                <p className="mt-6 text-lg leading-8 text-gray-600">
                  Meet the developers behind this CSUSM Capstone Project.
                </p>
              </div>
              <ul role="list" className="grid gap-x-8 gap-y-12 sm:grid-cols-2 sm:gap-y-16 xl:col-span-2">
                {people.map((person) => (
                  <li key={person.name}>
                    <div className="flex items-center gap-x-6">
                      <img className="h-28 w-28 rounded-full" src={person.imageUrl} alt="" />
                      <div>
                        <h3 className="text-base font-semibold leading-7 tracking-tight text-gray-900">{person.name}</h3>
                        <p className="text-sm font-semibold leading-6 text-indigo-600">{person.role}</p>
                      </div>
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </div>


      <button
        onClick={() => window.scrollTo({ top: isTop ? document.body.scrollHeight : 0, behavior: 'smooth' })}
        className="fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-black text-white w-12 h-12 rounded-full flex items-center justify-center text-2xl animate-bounce200"
      >
        {isTop ? '↓' : '↑'}
      </button>
    </div>
  )
}

function MyComponent() {
  return (
    <div>
      Hello, this is my component!
    </div>
  );
}