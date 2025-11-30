import React from 'react';
import FormInput from './FormInput';
import { Leaf, MapPin } from 'lucide-react';

function StepOne({ register, errors }) {
  return (
    <div className="grid gap-4">
      <FormInput label="Farm Title" name="title" icon={<Leaf />} register={register} errors={errors} required />
      <FormInput label="Crop Type" name="crop" icon={<Leaf />} register={register} errors={errors} required />
      <FormInput label="Location" name="location" icon={<MapPin />} register={register} errors={errors} required />
      <FormInput label="Price Per Unit (₦)" name="price_per_unit" type="number" register={register} errors={errors} required />
      <FormInput label="Expected ROI (e.g. 1.2 for 20%)" name="expected_roi" type="number" step="0.01" register={register} errors={errors} required />
    </div>
  );
}

export default StepOne;
