import {
  Show,
  SimpleShowLayout,
  TextField,
  useRecordContext,
} from "react-admin";

const CategoryTitle = () => {
  const record = useRecordContext();
  if (record == null) {
    return null;
  }
  return <div>Category {record.name}</div>;
};

export const CategoryShow = () => (
  <Show title={<CategoryTitle />}>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="name" />
    </SimpleShowLayout>
  </Show>
);
