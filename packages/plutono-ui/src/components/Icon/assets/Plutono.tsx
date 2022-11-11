import React, { FunctionComponent } from 'react';
import { SvgProps } from './types';

export const Plutono: FunctionComponent<SvgProps> = ({ size, ...rest }) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      x="0"
      y="0"
      enableBackground="new 0 0 351 365"
      version="1.1"
      viewBox="0 0 351 365"
      xmlSpace="preserve"
      height={size}
      width={size}
      {...rest}
    >
      <linearGradient x1="175.5" x2="175.5" y1="30%" y2="99%" gradientUnits="userSpaceOnUse">
        <stop offset="0" stopColor="#F05A28"></stop>
        <stop offset="1" stopColor="#FBCA0A"></stop>
      </linearGradient>
      <g
        aria-label="P"
        style={{ lineHeight: '1.25' }}
        fill="gray"
        strokeWidth="0.49"
        fontSize="19.604"
        transform="translate(-24.485 -8.048) scale(24.21752)"
      >
        <path
          fill="gray"
          d="M7.484.87q2.744 0 3.999 1.078 1.254 1.078 1.254 3.038 0 .863-.294 1.667-.274.784-.921 1.411-.647.628-1.725 1-1.078.353-2.647.353H5.543v5.45H3.778V.87zm-.157 1.51H5.543v5.527h1.411q1.333 0 2.215-.274.883-.294 1.314-.921.431-.628.431-1.647 0-1.353-.862-2.02-.863-.666-2.725-.666z"
        ></path>
      </g>
    </svg>
  );
};
