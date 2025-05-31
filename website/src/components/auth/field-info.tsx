import { AnyFieldApi } from '@tanstack/react-form/nextjs';

export function FieldInfo({ field }: { field: AnyFieldApi }) {
  return (
    <>
      {field.state.meta.isTouched && !field.state.meta.isValid ? (
        <em className='text-red-600'>
          {field.state.meta.errors.map(err => err.message).join(',')}
        </em>
      ) : null}
      {field.state.meta.isValidating ? 'Validating...' : null}
    </>
  );
}
