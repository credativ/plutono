import React from 'react';
import { render, screen } from '@testing-library/react';
import { ValiOptionFieldsProps, ValiOptionFields } from './ValiOptionFields';

const setup = (propOverrides?: ValiOptionFieldsProps) => {
  const queryType = 'range';
  const lineLimitValue = '1';
  const onLineLimitChange = jest.fn();
  const onQueryTypeChange = jest.fn();
  const onKeyDownFunc = jest.fn();

  const props: any = {
    queryType,
    lineLimitValue,
    onLineLimitChange,
    onQueryTypeChange,
    onKeyDownFunc,
  };

  Object.assign(props, propOverrides);

  return render(<ValiOptionFields {...props} />);
};

describe('ValiOptionFields', () => {
  it('should render step field', () => {
    setup();
    expect(screen.getByTestId('lineLimitField')).toBeInTheDocument();
  });

  it('should render query type field', () => {
    setup();
    expect(screen.getByTestId('queryTypeField')).toBeInTheDocument();
  });
});
