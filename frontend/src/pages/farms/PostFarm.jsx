import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useMutation } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


import { UploadToCloudinary } from '../../utilis/UpLoadToCloudinary';
import StepOne from '../../components/farms/StepOne';
import StepTwo from '../../components/farms/StepTwo';
import StepThree from '../../components/farms/StepThree'; // Assuming StepThree is defined similarly

const API_URL = 'http://localhost:4000/api/v1';

function PostFarm() {
  const [step, setStep] = useState(1);
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors, isSubmitting },
  } = useForm();

  const mutation = useMutation({
  mutationFn: async (formData) => {
    const file = formData.image[0];
    const imageUrl = await UploadToCloudinary(file);

    const payload = {
      ...formData,
      image_url: imageUrl,
      price_per_unit: Number(formData.price_per_unit),
      expected_roi: Number(formData.expected_roi),
      expected_yield: formData.expected_yield ? Number(formData.expected_yield) : undefined,
      expected_revenue: formData.expected_revenue ? Number(formData.expected_revenue) : undefined,
      production_duration: Number(formData.production_duration),
      units_available: Number(formData.units_available),
      start_date: formData.start_date,
    };

    const response = await axios.post(`${API_URL}/farms`, payload);
    return response.data.farm;
  },
    onSuccess: () => {
      toast.success('🎉 Farm posted successfully!');
      reset();
      setStep(1);
      navigate('/farms');
    },
    onError: (error) => {
      toast.error(`❌ ${error.response?.data?.message || 'Something went wrong'}`);
    },
  });

  const onSubmit = (data) => {
    mutation.mutate(data);
  };

  return (
    <div className='bg-[#F3E7D5] min-h-screen px-2 py-10 flex items-center justify-center'>
        <section className="bg-[#F3E7D5] min-h-screen px-4 py-10 flex items-center justify-center">
      <div className="backdrop-blur-md bg-white/20 border border-white/30 shadow-xl rounded-3xl p-8 max-w-3xl w-full">
        <h1 className="text-3xl font-bold text-[#1F3B17] mb-8 text-center">Post a New Farm</h1>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
          {step === 1 && <StepOne register={register} errors={errors} />}
          {step === 2 && <StepTwo register={register} errors={errors} />}
          {step === 3 && <StepThree register={register} errors={errors} />}

          <div className="flex justify-between items-center pt-4">
            {step > 1 && (
              <button
                type="button"
                onClick={() => setStep((s) => s - 1)}
                className="bg-[#5E7E3F] text-white px-6 py-2 rounded-lg hover:bg-[#2F5024]"
              >
                Back
              </button>
            )}
            {step < 3 ? (
              <button
                type="button"
                onClick={() => setStep((s) => s + 1)}
                className="bg-[#5E7E3F] text-white px-6 py-2 rounded-lg hover:bg-[#2F5024]"
              >
                Next
              </button>
            ) : (
              <button
                type="submit"
                disabled={isSubmitting || mutation.isPending}
                className="bg-[#5E7E3F] text-white px-6 py-2 rounded-lg hover:bg-[#2F5024]"
              >
                {mutation.isPending ? 'Submitting...' : 'Submit Farm'}
              </button>
            )}
          </div>
        </form>

       
      </div>
    </section>
    </div>
  );
}

export default PostFarm;
