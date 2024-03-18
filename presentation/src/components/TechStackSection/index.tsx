import githubWhiteIcon from '../../assets/svg/github-white.svg';
import reactWhteIcon from '../../assets/svg/react-logo-white.svg';
import tailwindWhiteIcon from '../../assets/svg/tailwind-logo-white.svg';
import dockerWhiteIcon from '../../assets/svg/docker-logo-white.svg';
import javacsriptWhiteIcon from '../../assets/svg/javascript-logo-white.svg';
import goWhiteIcon from '../../assets/svg/go-logo-white.svg';
import kubernetesWhiteIcon from '../../assets/svg/kubernetes-logo-white.svg';

import { Fade } from 'react-awesome-reveal';

export function TechStackSection() {
  return (
    <Fade>
      <div className="bg-gradient-to-r from-darkPink to-amaranthPink py-24 sm:py-32">
        <div className="mx-auto max-w-7xl px-6 lg:px-8">
          <h2 className="text-center text-lg font-semibold leading-8 text-white p-4">
            Built with the best tools in the industry
          </h2>
          <div className="mx-auto mt-10 grid max-w-lg grid-cols-7 items-center gap-x-8 gap-y-10 sm:max-w-xl sm:gap-x-10 lg:mx-0 lg:max-w-none">
            <img
              className="max-h-12 w-full object-contain lg:col-span-1"
              src={githubWhiteIcon}
              alt="Github"
              width={158}
              height={48}
            />
            <img
              className="max-h-12 w-full object-contain lg:col-span-1"
              src={reactWhteIcon}
              alt="React"
              width={158}
              height={48}
            />
            <img
              className="max-h-12 w-full object-contain lg:col-span-1"
              src={tailwindWhiteIcon}
              alt="Tailwind CSS"
              width={158}
              height={48}
            />
            <img
              className="max-h-12 w-full object-contain sm:col-start-2 lg:col-span-1"
              src={dockerWhiteIcon}
              alt="Docker"
              width={158}
              height={48}
            />
            <img
              className="max-h-12 w-full object-contain sm:col-start-auto lg:col-span-1"
              src={javacsriptWhiteIcon}
              alt="Javascript"
              width={158}
              height={48}
            />
            <img
              className="max-h-12 w-full object-contain sm:col-start-auto lg:col-span-1"
              src={goWhiteIcon}
              alt="GoLang"
              width={158}
              height={48}
            />
            <img
              className="max-h-12 w-full object-contain sm:col-start-auto lg:col-span-1"
              src={kubernetesWhiteIcon}
              alt="Kubernetes"
              width={158}
              height={48}
            />
          </div>
        </div>
      </div>
    </Fade>
  );
}
