import TextField from '@material-ui/core/TextField';

const TextInput = (props) => {
  return (
    <TextField
      fullWidth={props.fullWidth}
      label={props.label}
      placeholder={props.placeholder}
      helperText={props.helperText}
      margin='dense'
      multiline={props.multiline}
      required={props.required}
      rows={props.rows}
      value={props.value}
      type={props.type}
      onChange={props.onChange}
      variant={props.variant}
      InputProps={props.InputProps}
      InputLabelProps={{
        shrink: true,
      }}
    />
  );
};

export default TextInput;
