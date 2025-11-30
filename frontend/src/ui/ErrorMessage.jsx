// components/ui/ErrorMessage.jsx
import React from 'react';

export default function ErrorMessage({ message = 'Something went wrong' }) {
  return (
    <div className="text-center py-10 text-red-700">
      {message}
    </div>
  );
}
