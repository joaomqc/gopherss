import {
  Button,
  Datagrid,
  List,
  TextField,
  useRecordContext,
} from "react-admin";
import Mail from "@mui/icons-material/Mail";
import Drafts from "@mui/icons-material/DraftsOutlined";
import { useLocation } from "react-router-dom";

const ArticlePanel = () => {
  const record = useRecordContext();
  return <div>{record.content}</div>;
};

const ArticleReadButton = () => {
  const record = useRecordContext();
  if (record.read) {
    return (
      <Button
        onClick={(evt: any) => {
          evt.stopPropagation();
        }}
        startIcon={<Drafts />}
      />
    );
  }
  return (
    <Button
      onClick={(evt) => {
        evt.stopPropagation();
      }}
      startIcon={<Mail />}
    />
  );
};

export const ArticleList = () => {
  const filters: { [key: string]: any } = {};
  const location = useLocation();
  const search = location.search.slice(1);
  search.split("&").forEach((q) => {
    let qs = q.split("=");
    filters[qs[0]] = qs[1];
  });

  return (
    <List filter={filters}>
      <Datagrid rowClick="expand" expandSingle={true} expand={<ArticlePanel />}>
        <ArticleReadButton />
        <TextField
          source="title"
          fontStyle={() => {
            const record = useRecordContext();
            if (!record.read) {
              return { fontWeight: "bold" };
            }
          }}
        />
      </Datagrid>
    </List>
  );
};
