import { useState, useCallback, useEffect } from 'react';
import { TextInput, SelectBox } from '../UIkit';
import { fetchPrefMaster } from '../../reducks/utils/operations';
import { getPrefMaster } from '../../reducks/utils/selectors';
import { useDispatch, useSelector } from 'react-redux';

const RegisterAddress = () => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);

  const [name, setName]                   = useState(''),
    [postcode, setPostcode]               = useState(''),
    [prefOptions, setPrefOptions]         = useState(null),
    [prefId, setPrefId]                   = useState(0),
    [primaryAddrss, setPrimaryAddrss]     = useState(''),
    [secondaryAddrss, setSecondaryAddrss] = useState(''),
    [phoneNumber, setPhoneNumber]         = useState('');

  const inputName = useCallback((event) => {
    setName(event.target.value);
  }, [setName]);

  const inputPostcode = useCallback((event) => {
    setPostcode(event.target.value);
  }, [setPostcode]);

  const inputPrimaryAddress = useCallback((event) => {
    setPrimaryAddrss(event.target.value);
  }, [setPrimaryAddrss]);

  const inputSecondaryAddress = useCallback((event) => {
    setSecondaryAddrss(event.target.value);
  }, [setSecondaryAddrss]);

  const inputPhoneNumber = useCallback((event) => {
    setPhoneNumber(event.target.value);
  }, [setPhoneNumber]);

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
        placeholder={'山田 太郎'}
        helperText={'姓と名の間にスペースを入れてください'}
        multiline={false}
        required={true}
        rows={1}
        value={name}
        type={'text'}
        onChange={inputName}
      />
      <TextInput
        fullWidth={true}
        label={'郵便番号'}
        placeholder={'123-4567'}
        helperText={'ハイフンを入れてください'}
        multiline={false}
        required={true}
        rows={1}
        value={postcode}
        type={'text'}
        onChange={inputPostcode}
      />
      <div className='module-spacer--extra-small' />
      <SelectBox
        label={'都道府県'}
        required={true}
        value={prefId}
        options={prefOptions}
        select={setPrefId}
      />
      <TextInput
        fullWidth={true}
        label={'住所1'}
        placeholder={'市区町村 番地'}
        helperText={'ハイフンを入力してください'}
        multiline={false}
        required={true}
        rows={1}
        value={primaryAddrss}
        type={'text'}
        onChange={inputPrimaryAddress}
      />
      <TextInput
        fullWidth={true}
        label={'住所2'}
        placeholder={'建物名 部屋番号'}
        helperText={'マンションやアパートの方は必須です'}
        multiline={false}
        required={false}
        rows={1}
        value={secondaryAddrss}
        type={'text'}
        onChange={inputSecondaryAddress}
      />
      <TextInput
        fullWidth={true}
        label={'電話番号'}
        placeholder={'090-1234-5678'}
        helperText={'ハイフンを入れてください'}
        multiline={false}
        required={true}
        rows={1}
        value={phoneNumber}
        type={'text'}
        onChange={inputPhoneNumber}
      />
    </>
  );
};

export default RegisterAddress;
