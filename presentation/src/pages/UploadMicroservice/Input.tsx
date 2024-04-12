import React from 'react';
import PlusSvg from '../../assets/svg/plus.svg';
import MinusSvg from '../../assets/svg/minus.svg';
import { IInput } from '../../types/input';

interface IProps {
  index: number;
  input: IInput;
  inputsLength: number;
  handleInputChange: (index: number) => void;
  handleRemoveInput: (index: number) => void;
  handleAddInput: (index: number) => void;
}

export function Input(props: IProps) {
  return (
    <div className="flex items-center mx-4" key={props.index}>
      <div className="w-1/2">
        <div className="my-2 text-sm">Input Name</div>
        <input
          type="text"
          name={`Inputs[${props.index}].Name`}
          value={props.input.Name}
          onChange={() => props.handleInputChange(props.index)}
          className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
        />
      </div>
      <div className="w-1/2">
        <div className="my-2 text-sm">Input Data Type</div>
        <input
          type="text"
          name={`Inputs[${props.index}].DataType`}
          value={props.input.DataType}
          onChange={() => props.handleInputChange(props.index)}
          className="rounded-lg p-2 border w-2/3 border-gray-300 hover:shadow-md"
        />
        <button
          type="button"
          onClick={() => props.handleRemoveInput(props.index)}
          className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md"
        >
          <img src={MinusSvg} alt="remove" className="w-8 h-8" />
        </button>
        {props.index === props.inputsLength - 1 && (
          <button
            type="button"
            onClick={() => props.handleAddInput}
            className="bg-gray-300 rounded-lg py-1 px-2 hover:shadow-md"
          >
            <img src={PlusSvg} alt="add" className="w-8 h-8" />
          </button>
        )}
      </div>
    </div>
  );
}
