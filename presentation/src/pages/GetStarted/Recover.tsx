import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { IRecoveryData } from '../../types/recovery-data';
import { BackgroundImage } from '../../components/BackgroundImage';

const RecoverAccount = () => {
  const data = useRecovery();

  return (
    <BackgroundImage>
      <div className="container mx-auto p-4 lg:w-2/5">
        <div className="p-2 bg-gray-200 my-12 mx-auto rounded-xl drop-shadow-lg">
          <form onSubmit={data.handleSubmit}>
            {data.step === 'email' && (
              <>
                <div className="flex flex-col items-center mx-auto">
                  <div className="font-extrabold m-4 text-2xl">Account Recovery</div>
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
              </>
            )}
            {data.step === 'code' && (
              <>
                <div className="flex flex-col items-center mx-auto w-full">
                  <div className="my-2 text-sm">Verification Code</div>
                  <input
                    type="text"
                    name="Code"
                    value={data.formData.Code}
                    onChange={data.handleChange}
                    className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
                  />
                </div>
              </>
            )}
            {data.errorMessage && <div className="text-red-600 text-sm text-center my-4">{data.errorMessage}</div>}
            <div className="flex flex-col items-center mx-auto">
              <button className="bg-gray-800 text-white w-2/3 rounded-lg my-8 py-2 px-2 hover:shadow-md">
                {data.step === 'email' ? 'Send Reset Code' : 'Verify Code'}
              </button>
              <p className="text-sm mt-1">
                <span className="text-pink-500 hover:underline cursor-pointer" onClick={data.handleBackClick}>
                  Back to Login
                </span>
              </p>
            </div>
          </form>
        </div>
      </div>
    </BackgroundImage>
  );
};

function useRecovery() {
  const [formData, setFormData] = useState<IRecoveryData>({
    Email: '',
    Code: '',
  });

  const [step, setStep] = useState<'email' | 'code'>('email');
  const [errorMessage, setErrorMessage] = useState<string>('');

  const navigate = useNavigate();

  useEffect(() => {
    console.log(formData);
  }, [formData]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      let url = '';
      let body = {};

      if (step === 'email') {
        url = 'http://localhost:8080/recovery';
        body = { email: formData.Email };
      } else if (step === 'code') {
        url = 'http://localhost:8080/verify-code';
        body = { email: formData.Email, code: formData.Code };
      }

      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
      });

      if (response.ok) {
        const responseData = await response.json();
        console.log(responseData);
        if (step === 'email') {
          setStep('code');
        } else if (step === 'code') {
          navigate('/ResetPassword');
        }
      } else {
        const errorData = await response.json();
        console.error(errorData.error);
        setErrorMessage(errorData.error);
      }
    } catch (error) {
      console.error(error);
      setErrorMessage('Failed to recover account');
    }
  };

  const handleBackClick = () => {
    navigate('/developer-login');
  };

  return {
    step,
    formData,
    handleChange,
    handleSubmit,
    errorMessage,
    handleBackClick,
  };
}

export default RecoverAccount;
