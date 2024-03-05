import updateIcon from '../../assets/svg/update.svg';
import scaleIcon from '../../assets/svg/scale.svg';
import shieldIcon from '../../assets/svg/shield.svg';
import { Fade } from 'react-awesome-reveal';

export function AboutSection() {
  return (
    <Fade>
      <div id='about' className="relative isolate overflow-hidden bg-white px-6 py-24 sm:py-32 lg:overflow-visible lg:px-0">
        <div className="absolute inset-0 -z-10 overflow-hidden">
          <svg
            className="absolute left-[max(50%,25rem)] top-0 h-[64rem] w-[128rem] -translate-x-1/2 stroke-gray-200 [mask-image:radial-gradient(64rem_64rem_at_top,white,transparent)]"
            aria-hidden="true"
          >
            <defs>
              <pattern
                id="e813992c-7d03-4cc4-a2bd-151760b470a0"
                width={200}
                height={200}
                x="50%"
                y={-1}
                patternUnits="userSpaceOnUse"
              >
                <path d="M100 200V.5M.5 .5H200" fill="none" />
              </pattern>
            </defs>
            <svg x="50%" y={-1} className="overflow-visible fill-gray-50">
              <path
                d="M-100.5 0h201v201h-201Z M699.5 0h201v201h-201Z M499.5 400h201v201h-201Z M-300.5 600h201v201h-201Z"
                strokeWidth={0}
              />
            </svg>
            <rect width="100%" height="100%" strokeWidth={0} fill="url(#e813992c-7d03-4cc4-a2bd-151760b470a0)" />
          </svg>
        </div>
        <div className="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 lg:mx-0 lg:max-w-none lg:grid-cols-2 lg:items-start lg:gap-y-10">
          <div className="lg:col-span-2 lg:col-start-1 lg:row-start-1 lg:mx-auto lg:grid lg:w-full lg:max-w-7xl lg:grid-cols-2 lg:gap-x-8 lg:px-8">
            <div className="lg:pr-4">
              <div className="lg:max-w-lg">
                <p className="text-base font-semibold leading-7 text-selectiveYellow">Microservices</p>
                <h1 className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">Flexible, robust systems</h1>
                <p className="mt-6 text-xl leading-8 text-gray-700">
                  Microservices are a way of breaking up large applications into smaller,
                  more manageable pieces. Each piece is designed to be a self-contained
                  unit of functionality, which can be developed, tested, and deployed independently.
                  This makes it easier to understand, develop, and maintain complex systems,
                  and can also help to improve scalability and resilience.
                </p>
              </div>
            </div>
          </div>
          <div className="-ml-12 -mt-12 p-12 lg:sticky lg:top-4 lg:col-start-2 lg:row-span-2 lg:row-start-1 lg:overflow-hidden">
            <img
              className="w-full max-w-none rounded-xl bg-gray-900 shadow-xl ring-1 ring-gray-400/10 sm:w-full mb-4"
              src="https://imgix.datadoghq.com/img/knowledge-center/serverless-architecture/monolith-vs-microservices.png?auto=format&fit=max&w=847&dpr=2"
              alt=""
            />
            <img
              className="w-full max-w-none rounded-xl bg-gray-900 shadow-xl ring-1 ring-gray-400/10 sm:w-full"
              src="https://imgix.datadoghq.com/img/knowledge-center/serverless-architecture/sample-serverless-microservice.png?auto=format&fit=max&w=847&dpr=2"
              alt=""
            />
          </div>

          <div className="lg:col-span-2 lg:col-start-1 lg:row-start-2 lg:mx-auto lg:grid lg:w-full lg:max-w-7xl lg:grid-cols-2 lg:gap-x-8 lg:px-8">
            <div className="lg:pr-4">
              <div className="max-w-xl text-base leading-7 text-gray-700 lg:max-w-lg">
                <p>
                  The shift from monolithic architectures to microservices represents a significant
                  evolution in software development. Monoliths are characterized by their single,
                  unified codebase, which can simplify development and deployment in the early
                  stages but becomes increasingly cumbersome as applications grow.
                  Microservices, on the other hand, offer greater flexibility and
                  scalability by decomposing applications into smaller, independently deployable services.
                </p>
                <ul role="list" className="mt-8 space-y-8 text-gray-600">
                  <li className="flex gap-x-3">
                    <img src={scaleIcon} className="mt-1 h-5 w-5 flex-none text-indigo-600" aria-hidden="true" />
                    <span>
                      <strong className="font-semibold text-gray-900">Scalability.</strong> Automatically scales with demand, 
                      ensuring that applications can handle varying loads without manual intervention.
                    </span>
                  </li>
                  <li className="flex gap-x-3">
                    <img src={updateIcon} className="mt-1 h-5 w-5 flex-none text-indigo-600" aria-hidden="true" />
                    <span>
                      <strong className="font-semibold text-gray-900">Rapid Deployment and Updates.</strong> Allows for quicker 
                      updates and deployments, enabling a faster response to market changes or customer needs.
                    </span>
                  </li>
                  <li className="flex gap-x-3">
                    <img src={shieldIcon} className="mt-1 h-5 w-5 flex-none text-indigo-600" aria-hidden="true" />
                    <span>
                      <strong className="font-semibold text-gray-900">Resilience and Isolation.</strong> Failures in one service 
                      do not directly impact others, enhancing the overall stability of the application.
                    </span>
                  </li>
                </ul>
                <p className="mt-8">
                Serverless microservices offer a compelling model for developing modern, cloud-native applications. 
                By leveraging serverless computing, organizations can build applications that are more scalable, 
                resilient, and cost-effective. However, the adoption of this architecture requires careful 
                consideration of its challenges, particularly around service granularity, performance optimization, 
                and comprehensive monitoring. With the right approach and tools, serverless microservices can 
                significantly enhance an organization's ability to deliver innovative and responsive applications.
                </p>
              
              </div>
            </div>
          </div>
        </div>
      </div>
    </Fade>
  )
}
