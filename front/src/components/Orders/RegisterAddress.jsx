import { useState, useEffect } from 'react';
import { TextInput, SelectBox } from '../UIkit';
import { fetchPrefMaster } from '../../reducks/utils/operations';
import { getPrefMaster } from '../../reducks/utils/selectors';
import { useDispatch, useSelector } from 'react-redux';

const RegisterAddress = (props) => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);

  const inputName = (event) => {
    props.setAddress({ ...props.address, name: event.target.value });
  };
  const inputPostcode = (event) => {
    props.setAddress({ ...props.address, postcode: event.target.value });
  };
  const inputPrefId = (value) => {
    props.setAddress({ ...props.address, prefId: value });
  };
  const inputPrimaryAddress = (event) => {
    props.setAddress({ ...props.address, primaryAddress: event.target.value });
  };
  const inputSecondaryAddress = (event) => {
    props.setAddress({ ...props.address, secondaryAddress: event.target.value });
  };
  const inputPhoneNumber = (event) => {
    props.setAddress({ ...props.address, phoneNumber: event.target.value });
  };

  const [prefOptions, setPrefOptions] = useState(null);
  useEffect(() => {
    dispatch(fetchPrefMaster());
  }, []);
  const storePrefMaster = getPrefMaster(selector);
  useEffect(() => {
    if (storePrefMaster) {
      const tmpOptions = [{
        id: 0,
        name: '-',
      }];
      tmpOptions.push(...storePrefMaster);
      setPrefOptions(tmpOptions);
    }
  }, [storePrefMaster]);

  return (
    <>
      <h3 className='u-text__headline__k'>お届け先</h3>
      <TextInput
        fullWidth={true}
        label={'お名前'}
        helperText={'姓と名の間にスペースを入れてください（例：山田 太郎）'}
        multiline={false}
        required={true}
        rows={1}
        value={props.address.name}
        type={'text'}
        onChange={inputName}
      />
      <TextInput
        fullWidth={true}
        label={'郵便番号'}
        helperText={'ハイフンを入れてください（例：123-4567）'}
        multiline={false}
        required={true}
        rows={1}
        value={props.address.postcode}
        type={'text'}
        onChange={inputPostcode}
      />
      <div className='module-spacer--extra-small' />
      <SelectBox
        label={'都道府県'}
        required={true}
        value={props.address.prefId}
        options={prefOptions}
        select={inputPrefId}
      />
      <TextInput
        fullWidth={true}
        label={'住所1'}
        helperText={'例：新宿区西新宿2-8-1'}
        multiline={false}
        required={true}
        rows={1}
        value={props.address.primaryAddress}
        type={'text'}
        onChange={inputPrimaryAddress}
      />
      <TextInput
        fullWidth={true}
        label={'住所2'}
        helperText={'マンションやアパートの方は必須です（例：マンションボンゴA 201号室）'}
        multiline={false}
        required={false}
        rows={1}
        value={props.address.secondaryAddress}
        type={'text'}
        onChange={inputSecondaryAddress}
      />
      <TextInput
        fullWidth={true}
        label={'電話番号'}
        helperText={'ハイフンを入れてください（例：090-1234-5678）'}
        multiline={false}
        required={true}
        rows={1}
        value={props.address.phoneNumber}
        type={'text'}
        onChange={inputPhoneNumber}
      />
    </>
  );
};

export default RegisterAddress;
