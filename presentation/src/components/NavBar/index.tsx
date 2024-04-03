import { Link } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';
import { useEffect, useState } from 'react';

interface IProps {
    isDeveloper: boolean;
}
export function NavBar(props: IProps) {
    const data = useNavBar();
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
                {data.isAuthenticated &&(
                    <Link className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold" to="/Microservices">
                        Microservices
                    </Link>
                )}
                {data.isAuthenticated && data.userType === "Developer" && (
                    <Link className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold" to="/UploadMicroservice">
                        Upload
                    </Link>
                )}
                {!data.isAuthenticated && (
                    <Link className="hover:scale-105 transition duration-150 ease-in-out bg-darkPink text-white px-4 py-2 rounded" to="/developer-login">
                        Login
                    </Link>
                )}
                {data.isAuthenticated && (
                    <div className=" text-white font-bold">
                        Welcome {data.userName}!
                    </div>
                )}
                {data.isAuthenticated && (
                    <button className="hover:scale-105 transition duration-150 ease-in-out bg-darkPink text-white px-4 py-2 rounded" onClick={data.handleLogout}>
                        Logout
                    </button>
                )}
            </div>
        </div>
    );
}

function useNavBar(){
    const auth = useAuth()
    const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
    const [userType, setUserType] = useState<string | undefined>('');
    const [userName, setUserName] = useState<string | undefined>('');

    function handleLogout(){
        auth.logout();
        setIsAuthenticated(false);
    }

    useEffect(() => {
        //console.log(auth.isAuthenticated);
        auth.fetchUserDetails();
    }, []);

    useEffect(() => {
        if(auth.isAuthenticated){
            setIsAuthenticated(true);
            setUserType(auth.userDetails?.UserType);
            setUserName(auth.userDetails?.Username);
        }
    }, [auth.isAuthenticated]);

    return { isAuthenticated, userType, userName, handleLogout };
}