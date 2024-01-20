import { ReactNode, useState } from "react";
import { IoMdArrowDropdown } from "react-icons/io";

export type Option = {
  label: ReactNode;
  value: any;
  selected: boolean;
};

export type SelectProps = {
  options: Option[];
  setOption: (value: any) => void;
};

const Select = ({ options, setOption }: SelectProps) => {
  const [isOpen, setOpen] = useState(false);

  const toggleOptions = () => {
    setOpen((pre) => !pre);
  };

  return (
    <div className="relative">
      <div
        className="flex gap-3 items-center px-3 py-2 hover:text-sub cursor-pointer bg-zinc-600 rounded-md"
        onClick={toggleOptions}
      >
        {options.map((option) => {
          if (option.selected)
            return (
              <div className="flex items-center justify-center">
                {option.label}
                <IoMdArrowDropdown className="ml-5" />
              </div>
            );
        })}
      </div>

      <div
        className={`${
          isOpen ? "block" : "hidden"
        } absolute bg-zinc-600 rounded-md border-2 border-sub min-w-48`}
      >
        {options.map(({ label, value, selected }) => {
          return (
            <div
              className="px-6 py-3 cursor-pointer hover:bg-zinc-500 z-50"
              onClick={() => {
                if (selected) {
                  toggleOptions();
                  return;
                }

                setOption(value);
                toggleOptions();
              }}
            >
              {label}
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Select;
