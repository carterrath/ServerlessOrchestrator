import { Fade } from 'react-awesome-reveal';
import photoMae from '../../assets/images/photoMae.png';
import photoRuth from '../../assets/images/photoRuth.png';
import photoCarter from '../../assets/images/photoCarter.png';
import photoJackie from '../../assets/images/photoJackie.png';
import linkedInSvg from '../../assets/svg/linkedIn.svg';
import githubBlackSvg from '../../assets/svg/github-black.svg';

export function TeamSection() {
  const people = [
    {
      name: 'Mae Pereyra',
      role: 'Software Barbie',
      imageUrl: photoMae,
      linkedinUrl: 'https://www.linkedin.com/in/macariamae/',
      githubUrl: 'https://github.com/maepereyra'
    },
    {
      name: 'Ruth Jimenez',
      role: 'Software Barbie',
      imageUrl: photoRuth,
      linkedinUrl: 'https://www.linkedin.com/in/ruth-jimenez-4826651a8/',
      githubUrl: 'https://github.com/ruthijimenez'
    },
    {
      name: 'Carter Rath',
      role: 'Software Barbie',
      imageUrl: photoCarter,
      linkedinUrl: 'https://www.linkedin.com/in/carter-rath/',
      githubUrl: 'https://github.com/carterrath'
    },
    {
      name: 'Jaclyn Walsh',
      role: 'Software Barbie',
      imageUrl: photoJackie,
      linkedinUrl: 'https://www.linkedin.com/in/jaclynewalsh/',
      githubUrl: 'https://github.com/JaclynW'
    },
  ];

  return (
    <Fade>
      <div className="border-b">
        <div className="flex justify-center">
          <div className="bg-white p-6 w-full max-w-7xl">
            <div className="py-24 sm:py-32">
              <div className="mx-auto grid gap-x-8 gap-y-20 px-6 lg:px-8 xl:grid-cols-3">
                <div className="max-w-2xl">
                  <h2 className="text-3xl font-bold leading-9 tracking-tighter sm:text-3xl md:text-6xl">
                    Meet our team
                  </h2>
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
                          <h3 className="text-base font-semibold leading-7 tracking-tight text-gray-900">
                            {person.name}
                          </h3>
                          <p className="text-sm font-semibold leading-6 text-indigo-600">{person.role}</p>
                          <div className="flex gap-2 mt-2">
                            <a href={person.linkedinUrl} target="_blank" rel="noreferrer">
                              <img src={linkedInSvg} alt="linkedin" className="w-6 h-6 hover:scale-[110%] transition duration-150 ease-in-out" />
                            </a>
                            <a href={person.githubUrl} target="_blank" rel="noreferrer">
                              <img src={githubBlackSvg} alt="github" className="w-6 h-6 hover:scale-[110%] transition duration-150 ease-in-out" />
                            </a>
                          </div>
                        </div>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Fade>
  );
}
