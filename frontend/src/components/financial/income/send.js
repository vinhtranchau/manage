import React, { useReducer, useEffect, useContext, useState } from "react";
import { Params } from "../../../proto/financial/financial_pb";
import Context from "../../context/context";
import { toDestination } from "../../../utils/backend";

export default function SendTo(props) {
  const context = useContext(Context);
  const client = context.finances;

  const [incomes, setIncomes] = useState([]);

  const [userInput, setUserInput] = useReducer(
    (state, newState) => ({ ...state, ...newState }),
    {
      name: "",
      amount: 0,
      to: "",
      max: 0
    }
  );

  const handleChange = (name, value) => {
    setUserInput({ [name]: value });
  };

  function submit() {
    let params = new Params();

    params.setName(userInput.name);
    params.setAmount(userInput.amount);
    params.setTo(userInput.to);

    toDestination(params, (err, res) => {
      if (err) {
        console.log(err);
      }

      if (res) {
        props.GetBanks();
      }
    });
  }

  useEffect(() => {
    handleChange("max", incomes[userInput.name]);
    //eslint-disable-next-line react-hooks/exhaustive-deps
  }, [userInput.name]);

  const Incomes = () => {
    let result = {};

    let data = {};
    for (var income of props.incomes) {
      for (let sector of income.sectorsList) {
        if (data[sector.name]) {
          data[sector.name] += sector.amount;
        } else {
          data[sector.name] = sector.amount;
        }
      }
    }

    for (let key of Object.keys(data)) {
      result[key] = data[key];
    }

    setIncomes(result);
  };

  useEffect(() => {
    Incomes();
  }, [props.incomes]);

  return (
    <div className="d-flex flex-column justify-content-center ">
      <h6 className="mx-auto mb-3">Send income to destination</h6>

      <select
        className="custom-select income-input bg-success"
        onChange={e => handleChange("name", e.target.value)}
      >
        <option value="">Select a source from</option>
        {Object.keys(incomes).map((income, i) => (
          <option key={i}>{income}</option>
        ))}
      </select>

      <select
        className="custom-select bg-success income-input"
        onChange={e => handleChange("to", e.target.value)}
      >
        <option value="">Select a destination</option>
        {props.banks.map((bank, i) => (
          <option key={i}>{bank.name}</option>
        ))}
      </select>

      <input
        className="form-control w-75 mt-4 mx-auto"
        type="number"
        value={userInput.amount}
        min="0"
        max={userInput.max}
        onChange={e => {
          if (e.target.value <= userInput.max) {
            handleChange("amount", e.target.value);
          }
        }}
        placeholder="R$"
      />
      <div className="w-75 mt-4 mx-auto">
        <button
          onClick={() => submit()}
          className="btn btn-success float-right"
        >
          Send
        </button>
      </div>
    </div>
  );
}
