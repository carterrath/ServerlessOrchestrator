import { Link } from 'react-router-dom';

interface IProps {
    isDeveloper: boolean;
}
export function NavBar(props: IProps) {
    return (
        <div className="top-0 z-[999999999999] sticky w-full h-20 px-4 items-center flex drop-shadow-xl bg-gradient-to-r from-darkPink to-amaranthPink justify-between">
            <div className="flex gap-6 items-center">
                <Link className="flex items-center hover:scale-105 transition duration-150 ease-in-out text-xl" to="/Home">
                    <img className="w-8 h-8" src="/src/assets/images/logo.png" alt="Logo" />
                    &nbsp;<b className="text-white">Serverless Orcastrator</b>
                </Link>

            </div>
            <div className="flex gap-8 items-center">
                <Link className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold" to="/Home">
                    Home
                </Link>
                <Link className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold" to="/Microservices">
                    Microservices
                </Link>
                {props.isDeveloper && (
                    <Link className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold" to="/UploadMicroservice">
                        Upload
                    </Link>
                )}
                <Link className="hover:scale-105 transition duration-150 ease-in-out bg-darkPink text-white px-4 py-2 rounded" to="/developer-login">
                    Login
                </Link>
            </div>
        </div>
    );
}