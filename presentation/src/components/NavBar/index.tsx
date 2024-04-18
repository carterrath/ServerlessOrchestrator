import { Link } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';

export function NavBar() {
  const data = useNavBar();
  return (
    <div className="top-0 z-[999999999999] sticky w-full h-20 px-4 items-center flex drop-shadow-xl bg-gradient-to-r from-darkPink to-amaranthPink justify-between">
      <div className="flex gap-6 items-center">
        <Link className="flex items-center hover:scale-105 transition duration-150 ease-in-out text-xl" to="/Home">
          <img className="w-8 h-8" src="/src/assets/images/logo.png" alt="Logo" />
          &nbsp;<b className="text-white">Serverless Orchestrator</b>
        </Link>
      </div>
      <div className="flex gap-8 items-center">
        <Link className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold" to="/Home">
          Home
        </Link>
        {data.isAuthenticated && (
          <Link
            className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold"
            to="/Microservices"
          >
            Microservices
          </Link>
        )}
        {data.isAuthenticated && data.userType === 'Developer' && (
          <Link
            className="hover:scale-105 transition duration-150 ease-in-out text-white font-bold"
            to="/UploadMicroservice"
          >
            Upload
          </Link>
        )}
        {!data.isAuthenticated && (
          <Link
            className="hover:scale-105 transition duration-150 ease-in-out bg-darkPink text-white px-4 py-2 rounded"
            to="/DeveloperLogin"
          >
            Login
          </Link>
        )}
        {data.isAuthenticated && <div className=" text-white font-bold">Welcome {data.userName}!</div>}
        {data.isAuthenticated && (
          <button
            className="hover:scale-105 transition duration-150 ease-in-out bg-darkPink text-white px-4 py-2 rounded"
            onClick={data.handleLogout}
          >
            Logout
          </button>
        )}
      </div>
    </div>
  );
}

function useNavBar() {
  const auth = useAuth();

  const isAuthenticated = auth?.isAuthenticated;
  const userType = auth?.userDetails?.UserType;
  const userName = auth?.userDetails?.Username;

  function handleLogout() {
    auth?.logout();
  }

  return { isAuthenticated, userType, userName, handleLogout };
}
