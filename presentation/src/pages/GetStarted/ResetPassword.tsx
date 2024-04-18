// this will prompt user to enter a new password
import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { IResetData } from '../../types/password-reset';
import { BackgroundImage } from '../../components/BackgroundImage';

const ResetPassword = () => {
  const data = useReset();

  return (
    <BackgroundImage>
      <div className="container mx-auto p-4 lg:w-2/5">
        <div className="p-2 bg-gray-200 my-12 mx-auto rounded-xl drop-shadow-lg">
          <form onSubmit={data.handleSubmit}>
            <div className="flex flex-col items-center mx-auto">
              <div className="font-extrabold m-4 text-2xl">Reset Password</div>
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
              <div className="my-2 text-sm w-2/3 text-left ml-2 font-semibold">New Password</div>
              <input
                type="password"
                name="Password"
                value={data.formData.Password}
                onChange={data.handleChange}
                className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
              />
            </div>
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-800 text-white w-2/3 rounded-lg my-8 py-2 px-2 hover:shadow-md">Reset</button>
            </div>
          </form>
        </div>
      </div>
    </BackgroundImage>
  );
};

function useReset() {
  const navigate = useNavigate();

  const [formData, setFormData] = useState<IResetData>({
    Email: '',
    Password: '',
  });

  useEffect(() => {
    console.log(formData);
  }, [formData]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log(e.target.name, e.target.value);
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/reset', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        const responseData = await response.json();
        console.log(responseData);
        navigate('/Home'); // Navigate on success
      } else {
        // Log the response status
        console.error('Failed to reset password. Status:', response.status);
        // Handle HTTP error responses (e.g., 400, 401, 500)
        // Optionally, log the response body
        const errorData = await response.json();
        console.error('Error response:', errorData);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return {
    formData,
    handleChange,
    handleSubmit,
  };
}

export default ResetPassword;
