create table pdvs (
	id_pdvs bigserial,
	tenant varchar,
	nome varchar,
	cidade varchar,
	endereco varchar,
	cep varchar
);

create index in_pdvs_tenant on pdvs (tenant, id_pdvs);
