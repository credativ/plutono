import React, { FC } from 'react';
import { OrgRole } from '@credativ/plutono-data';
import { Select } from '@credativ/plutono-ui';

interface Props {
  value: OrgRole;
  onChange: (role: OrgRole) => void;
}

const options = Object.keys(OrgRole).map((key) => ({ label: key, value: key }));

export const OrgRolePicker: FC<Props> = ({ value, onChange, ...restProps }) => (
  <Select
    value={value}
    options={options}
    onChange={(val) => onChange(val.value as OrgRole)}
    placeholder="Choose role..."
    {...restProps}
  />
);
