import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { IUserData } from '../../types/user-data';

const DeveloperLogin = () => {
  const data = useDevLogin();
  
  return (
    <div className="container mx-auto p-4 lg:w-2/5">
      <div className="p-2 bg-gray-200 my-12 mx-auto rounded-xl drop-shadow-lg">
          <form onSubmit={data.handleSubmit}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">
                Developer Login!
              </div>
            </div>
            <div className="flex flex-col items-center mx-auto w-full">
              <div className="my-2 text-sm">
                Username
              </div>
              <input
                type="text"
                name="Username"
                value={data.formData.Username}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="flex flex-col items-center mx-auto w-full">
              <div className="my-2 text-sm">
                Password
              </div>
              <input
                type="password"
                name="Password"
                value={data.formData.Password}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-300 rounded-lg m-4 py-2 px-2 hover:shadow-md">
                Login
              </button>
              <p className="text-sm mt-1">
                Don't have an account?
                <span 
                  className="text-pink-500 hover:underline cursor-pointer" 
                  onClick={data.handleSignupClick}>
                  Signup
                </span>
              </p>
              <p className="text-sm mt-1">
                Are you a Comsumer?
                <span 
                  className="text-pink-500 hover:underline cursor-pointer" 
                  onClick={data.handleConLoginClick}>
                  Comsumer Login
                </span>
              </p>
            </div>
          </form>
        </div>
    </div>
  );
};

function useDevLogin() {
  const navigate = useNavigate();

  const [formData, setFormData] = useState<IUserData>({
    Email: '',
    Username: '',
    Password: '',
    UserType: 'Developer',
  });

  useEffect(() => {
    console.log(formData);
  }, [formData]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log(e.target.name, e.target.value); // Add this line
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/login/developer', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        // body: JSON.stringify(formData),
        body: JSON.stringify({
          username: formData.Username,
          password: formData.Password,
          userType: 'Developer'
        }),
      });
      
      if (response.ok) {
        const responseData = await response.json();
        console.log(responseData);
        navigate('/Microservices'); // Navigate on success
      } else {
          const errorData = await response.json();
          console.error(errorData.error);
        // Handle HTTP error responses (e.g., 400, 401, 500)
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleSignupClick = () => {
    navigate('/developer-signup');
  };

  const handleConLoginClick = () => {
    navigate('/consumer-login');
  };

  return{
    formData,
    handleChange,
    handleSubmit,
    handleSignupClick,
    handleConLoginClick,
  }
}

export default DeveloperLogin;
