import photoMae from '../../assets/images/photoMae.png';
import photoRuth from '../../assets/images/photoRuth.png';
import photoCarter from '../../assets/images/photoCarter.png';
import photoJackie from '../../assets/images/photoJackie.png';

export function TeamSection() {
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

    return (
        <div className="border-b border-slate-200">
        <div className="flex justify-center">
            <div className="bg-white p-6 w-full max-w-7xl">
                <div className="py-24 sm:py-32">
                    <div className="mx-auto grid gap-x-8 gap-y-20 px-6 lg:px-8 xl:grid-cols-3">
                        <div className="max-w-2xl">
                            <h2 className="text-3xl font-bold leading-9 tracking-tighter sm:text-3xl md:text-6xl">Meet our team</h2>
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
        </div>
    );
}