import React from 'react';

function FormInput({ label, name, register, errors, type = 'text', step, required, icon }) {
  const validationOptions = {
    required,
    ...(type === 'number' && { valueAsNumber: true }),
  };

  return (
    <div>
      <label htmlFor={name} className="block text-sm font-medium text-[#1F3B17] mb-1 flex items-center gap-2">
        {icon} {label}
      </label>
      <input
        id={name}
        type={type}
        step={step}
        {...register(name, validationOptions)}
        className={`w-full border rounded-md px-3 py-2 focus:outline-none focus:ring-2 backdrop-blur-md bg-white/30 ${
          errors[name] ? 'border-red-500 ring-red-300' : 'border-gray-300 focus:ring-[#5E7E3F]'
        }`}
      />
      {errors[name] && <p className="text-red-600 text-xs mt-1">This field is required</p>}
    </div>
  );
}

export default FormInput;
