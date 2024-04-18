import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { IUser } from '../../types/user-upload';
import { BackgroundImage } from '../../components/BackgroundImage';

const ConsumerSignup = () => {
  const data = useConSignup();

  return (
    <BackgroundImage>
      <div className="container mx-auto p-4 lg:w-2/5">
        <div className="p-2 bg-gray-200 my-12 mx-auto rounded-xl drop-shadow-lg">
          <form onSubmit={data.handleSubmit}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">Consumer Signup!</div>
            </div>
            <div className="flex flex-col items-center mx-auto w-full">
              <div className="my-2 text-sm w-2/3 text-left ml-2 font-semibold">Email</div>
              <input
                type="email"
                name="Email"
                value={data.formData.Email}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="flex flex-col items-center mx-auto w-full">
              <div className="my-2 text-sm w-2/3 text-left ml-2 font-semibold">Username</div>
              <input
                type="text"
                name="Username"
                value={data.formData.Username}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="flex flex-col items-center mx-auto w-full">
              <div className="my-2 text-sm w-2/3 text-left ml-2 font-semibold">Password</div>
              <input
                type="password"
                name="Password"
                value={data.formData.Password}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-800 text-white w-2/3 rounded-lg my-8 py-2 px-2 hover:shadow-md">Signup</button>
              <p className="text-sm mt-1">
                Already have an account?&nbsp;
                <span className="text-pink-500 hover:underline cursor-pointer" onClick={data.handleLoginClick}>
                  Login
                </span>
              </p>
              <p className="text-sm mt-1 mb-2">
                Are you a Developer?&nbsp;
                <span className="text-pink-500 hover:underline cursor-pointer" onClick={data.handleDevSignupClick}>
                  Developer Signup
                </span>
              </p>
            </div>
          </form>
        </div>
      </div>
    </BackgroundImage>
  );
};

function useConSignup() {
  const navigate = useNavigate();

  const [formData, setFormData] = useState<IUser>({
    Email: '',
    Username: '',
    Password: '',
    UserType: 'Consumer',
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
      const response = await fetch('http://localhost:8080/signup/consumer', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        const responseData = await response.json();
        console.log(responseData);
        navigate('/ConsumerLogin');
      } else {
        console.error('Failed to signup');
        // Handle HTTP error responses (e.g., 400, 401, 500)
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleLoginClick = () => {
    navigate('/ConsumerLogin'); // need to create this page
  };

  const handleDevSignupClick = () => {
    navigate('/DeveloperSignup'); // need to create this page
  };

  return {
    formData,
    handleChange,
    handleSubmit,
    handleLoginClick,
    handleDevSignupClick,
  };
}

export default ConsumerSignup;
