// StepThree.jsx
import React from 'react';
import { ImageIcon } from 'lucide-react';

function StepThree({ register, errors }) {
  return (
    <div className="grid gap-4">
      <div>
        <label htmlFor="image" className="block text-sm font-medium text-[#1F3B17] mb-1 flex items-center gap-2">
          <ImageIcon /> Upload Image
        </label>
        <input
          type="file"
          id="image"
          accept="image/*"
          {...register('image', { required: true })}
          className={`w-full border rounded-md px-3 py-2 focus:outline-none focus:ring-2 backdrop-blur-md bg-white/30 ${
            errors.image ? 'border-red-500 ring-red-300' : 'border-gray-300 focus:ring-[#5E7E3F]'
          }`}
        />
        {errors.image && <p className="text-red-600 text-xs mt-1">This field is required</p>}
      </div>

      <div>
        <label htmlFor="description" className="block text-sm font-medium text-[#1F3B17] mb-1">
          Farm Description
        </label>
        <textarea
          id="description"
          {...register('description', { required: true })}
          rows="4"
          placeholder="Write details about the farm..."
          className={`w-full border rounded-md p-3 focus:outline-none focus:ring-2 backdrop-blur-md bg-white/30 ${
            errors.description ? 'border-red-500 ring-red-300' : 'border-gray-300 focus:ring-[#5E7E3F]'
          }`}
        />
        {errors.description && <p className="text-red-600 text-xs mt-1">This field is required</p>}
      </div>
    </div>
  );
}

export default StepThree;
