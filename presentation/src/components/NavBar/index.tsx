import { Link } from 'react-router-dom';

interface IProps{
    isDeveloper: boolean;
}
export function NavBar(props: IProps){
    return(
        <div className="top-0 z-[999999999999] sticky w-full h-20 px-4 items-center flex drop-shadow-xl bg-gradient-to-r from-darkPink to-amaranthPink justify-between">
            <div className="flex gap-4 items-center">
                <Link className="flex items-center hover:scale-105 transition duration-150 ease-in-out" to="/Home">
                    <img className="w-8 h-8" src="/src/assets/images/logo.png" alt="Logo"/> 
                     &nbsp;<b>Serverless Orcastrator</b>
                </Link>
                <Link className="hover:scale-105 transition duration-150 ease-in-out" to="/Home">
                    Home
                </Link>
                <Link className="hover:scale-105 transition duration-150 ease-in-out" to="/Microservices">
                    Microservices
                </Link>
                {props.isDeveloper && (
                    <Link className="hover:scale-105 transition duration-150 ease-in-out" to="/UploadMicroservice">
                        Upload
                    </Link>
                )}     
            </div>
            <div className="flex gap-4">
            <Link className="hover:scale-105 transition duration-150 ease-in-out bg-white text-black px-4 py-2 rounded" to="/GetStarted">
                    Login
                </Link>
            </div>
        </div>
    );
}