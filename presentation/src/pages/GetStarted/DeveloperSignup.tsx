import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { IUser } from '../../types/user-upload';
import { BackgroundImage } from '../../components/BackgroundImage';
import { API_URL } from '../../constants';

const DeveloperSignup = () => {
  const data = useDevSignup();

  return (
    <BackgroundImage>
      <div className="container mx-auto p-4 lg:w-2/5">
        <div className="p-2 bg-gray-200 my-12 mx-auto rounded-xl drop-shadow-lg">
          <form onSubmit={data.handleSubmit}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">Developer Signup!</div>
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
              <button className="bg-gray-800 text-white rounded-lg w-2/3 my-8 py-2 px-2 hover:shadow-md">Signup</button>
              <p className="text-sm mt-1">
                Already have an account?&nbsp;
                <span className="text-pink-500 hover:underline cursor-pointer" onClick={data.handleLoginClick}>
                  Login
                </span>
              </p>
              <p className="text-sm mt-1 mb-3">
                Are you a Consumer?&nbsp;
                <span className="text-pink-500 hover:underline cursor-pointer" onClick={data.handleConSignupClick}>
                  Consumer Signup
                </span>
              </p>
            </div>
          </form>
        </div>
      </div>
    </BackgroundImage>
  );
};

function useDevSignup() {
  const navigate = useNavigate();

  const [formData, setFormData] = useState<IUser>({
    Email: '',
    Username: '',
    Password: '',
    UserType: 'Developer',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch(`${API_URL}/signup/developer`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        navigate('/DeveloperLogin');
      } else {
        console.error('Failed to signup');
        // Handle HTTP error responses (e.g., 400, 401, 500)
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleLoginClick = () => {
    navigate('/DeveloperLogin'); // need to create this page
  };

  const handleConSignupClick = () => {
    navigate('/ConsumerSignup'); // need to create this page
  };

  return {
    formData,
    handleChange,
    handleSubmit,
    handleLoginClick,
    handleConSignupClick,
  };
}

export default DeveloperSignup;
