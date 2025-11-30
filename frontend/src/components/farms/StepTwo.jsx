import React from 'react';
import FormInput from './FormInput';
import { CalendarDays } from 'lucide-react';

function StepTwo({ register, errors }) {
  return (
    <div className="grid gap-4">
      <FormInput label="Expected Yield (kg)" name="expected_yield" type="number" register={register} errors={errors} />
      <FormInput label="Expected Revenue (₦)" name="expected_revenue" type="number" register={register} errors={errors} />
      <FormInput label="Production Duration (days)" name="production_duration" type="number" register={register} errors={errors} required />
      <FormInput label="Units Available" name="units_available" type="number" register={register} errors={errors} required />
      <FormInput label="Start Date" name="start_date" type="date" icon={<CalendarDays />} register={register} errors={errors} required />
    </div>
  );
}

export default StepTwo;
