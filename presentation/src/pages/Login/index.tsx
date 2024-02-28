import { useNavigate } from "react-router-dom";
import { useState } from "react";

export function Login() {

  const data = useLogin();
    return (
    <>
    <div className="container mx-auto p-4 lg:w-2/5">
      <div className="p-2 bg-gray-100 my-12 mx-auto rounded-xl drop-shadow-lg">

        {/* Developer or Consumer form */}
        <div className="flex items-center justify-center mb-4">
            <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md" onClick={data.handleDeveloperSignup}>Developer</button>
            <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md" onClick={data.handleConsumerSignup}>Consumer</button>
          </div>

        {/* Developer signup form here */}
        {data.showDeveloperSignupForm && (
          <form onSubmit={data.handleLogin}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">
                Developer Signup!
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Email
                </div>
                <input
                  type="email"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Username
                </div>
                <input
                  type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Password
                </div>
                <input
                  type="password"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
              <div className="w-full">         
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="flex items-center bg-white -300 rounded-lg m-4 py-2 px-8 hover:shadow-md">
                <img className="w-8 h-8" src="/src/assets/images/google.png" alt="Logo"/> 
              Sign-In with Google
              </button>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className=" bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
                Signup
              </button>
              <p className="text-sm mt-1">
                Already have an account?
                <span 
                  className="text-pink-500 hover:underline cursor-pointer" 
                  onClick={data.showDeveloperLoginForm ? data.handleConsumerLogin : data.handleDeveloperLogin}>
                  Login
                </span>
              </p>
            </div>
          </form>
        )}


        {/* Consumer signup form here */}
        {data.showConsumerSignupForm && (
          <form onSubmit={data.handleLogin}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">
                Consumer Signup!
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Email
                </div>
                <input
                 type="email"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Username
                </div>
                <input
                 type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Password
                </div>
                <input
                  type="password"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
              <div className="w-full">         
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="flex items-center bg-white -300 rounded-lg m-4 py-2 px-8 hover:shadow-md">
                <img className="w-8 h-8" src="/src/assets/images/google.png" alt="Logo"/> 
              Sign-In with Google
              </button>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
                Signup
              </button>
              <p className="text-sm mt-1">
                Already have an account?
                <span 
                  className="text-pink-500 hover:underline cursor-pointer" 
                  onClick={data.showConsumerLoginForm ? data.handleDeveloperLogin : data.handleConsumerLogin}>
                  Login
                </span>
              </p>
            </div>
          </form>
        )}

        {/* Developer Login */}
        {data.showDeveloperLoginForm && (
          <form onSubmit={data.handleLogin}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">
                Developer
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Username
                </div>
                <input
                 type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Password
                </div>
                <input
                  type="password"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
              <div className="w-full">         
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="flex items-center bg-white -300 rounded-lg m-4 py-2 px-8 hover:shadow-md">
                <img className="w-8 h-8" src="/src/assets/images/google.png" alt="Logo"/> 
              Login with Google
              </button>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
                Login
              </button>
              <p className="text-sm mt-1">
                Don't have an account?
                <span 
                  className="text-pink-500 hover:underline cursor-pointer" 
                  onClick={data.showDeveloperSignupForm ? data.handleConsumerSignup : data.handleDeveloperSignup}>
                  Signup
                </span>
              </p>
            </div>
          </form>
        )}

        {/* Consumer Login Form */}
        {data.showConsumerLoginForm && (
          <form onSubmit={data.handleLogin}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">
                Consumer
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Username
                </div>
                <input
                  type="text"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <div className="flex flex-col items-center mx-auto w-full">
                <div className="my-2 text-sm">
                  Password
                </div>
                <input
                  type="password"
                  placeholder=""
                  className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"/>
              </div>
              <div className="w-full">         
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="flex items-center bg-white -300 rounded-lg m-4 py-2 px-8 hover:shadow-md">
                <img className="w-8 h-8" src="/src/assets/images/google.png" alt="Logo"/> 
              Login with Google
              </button>
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
                Login
              </button>
              <p className="text-sm mt-1">
                Don't have an account?
                <span 
                  className="text-pink-500 hover:underline cursor-pointer" 
                  onClick={data.showConsumerSignupForm ? data.handleDeveloperSignup : data.handleConsumerSignup}>
                  Signup
                </span>
              </p>
            </div>
          </form>
        )}
      </div>
    </div>
    </>
  );
}

function useLogin(){
  const navigate = useNavigate();

  const [showDeveloperLoginForm, setShowDeveloperLoginForm] = useState(false);
  const [showConsumerLoginForm, setShowConsumerLoginForm] = useState(false);
  const [showDeveloperSignupForm, setShowDeveloperSignupForm] = useState(false);
  const [showConsumerSignupForm, setShowConsumerSignupForm] = useState(false);

  const handleDeveloperSignup = () => {
    setShowDeveloperSignupForm(true);
    setShowConsumerSignupForm(false);
    setShowDeveloperLoginForm(false);
    setShowConsumerLoginForm(false);
  };

  const handleConsumerSignup = () => {
    setShowConsumerSignupForm(true);
    setShowDeveloperSignupForm(false);
    setShowDeveloperLoginForm(false);
    setShowConsumerLoginForm(false);
  };

  const handleDeveloperLogin = () => {
    setShowDeveloperLoginForm(true);
    setShowConsumerLoginForm(false);
    setShowDeveloperSignupForm(false);
    setShowConsumerSignupForm(false);
  };

  const handleConsumerLogin = () => {
    setShowConsumerLoginForm(true);
    setShowDeveloperLoginForm(false);
    setShowDeveloperSignupForm(false);
    setShowConsumerSignupForm(false);
  };

  const handleLogin = () => {
    navigate('/register');
  };

  return {
    handleLogin,
    handleDeveloperSignup,
    handleConsumerSignup,
    handleDeveloperLogin,
    handleConsumerLogin,
    showDeveloperSignupForm,
    showConsumerSignupForm,
    showDeveloperLoginForm,
    showConsumerLoginForm
  }
}