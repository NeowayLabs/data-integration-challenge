package pdv.l3ug1m.github.com.pdvservice.services;

import static pdv.l3ug1m.github.com.pdvservice.jooq.Tables.PDVS;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Future;

import org.apache.commons.lang3.ObjectUtils;
import org.apache.commons.lang3.StringUtils;
import org.jooq.DSLContext;
import org.jooq.Record4;
import org.jooq.RecordMapper;
import org.jooq.SelectConditionStep;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.AsyncResult;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import pdv.l3ug1m.github.com.pdvservice.jooq.tables.records.PdvsRecord;

@Service
public class PdvService {

	@Autowired
	DSLContext dslContext;

	@Transactional
	@Async
	public Future<Boolean> processLines(String tenantId, List<String[]> lines) {

		List<PdvsRecord> records = new ArrayList<>();

		for (String[] line : lines) {
			PdvBean bean = mappging(line);

			PdvsRecord record = new PdvsRecord();
			record.setTenant(tenantId);
			record.setNome(bean.nome);
			record.setCidade(bean.cidade);
			record.setEndereco(bean.endereco);
			record.setCep(bean.cep);

			records.add(record);
		}

		dslContext.batchInsert(records).execute();

		return new AsyncResult<>(true);
	}

	private PdvBean mappging(String[] line) {
		PdvBean bean = new PdvBean();
		bean.nome = ObjectUtils.defaultIfNull(StringUtils.trimToNull(line[0]), "NOME DESCONHECIDO");
		bean.cidade = ObjectUtils.defaultIfNull(StringUtils.trimToNull(line[1]), "NOME DESCONHECIDO");
		bean.endereco = line[2];
		bean.cep = line[3];

		bean.nome = bean.nome.replaceAll("\"", "");
		bean.cidade = bean.cidade.replaceAll("\"", "");
		bean.endereco = bean.endereco.replaceAll("\"", "");
		
		if (bean.cep.length() != 8) {
			bean.cep = "00000000";
		}
		

		return bean;
	}

	public List<PdvBean> list(String tenantId, String nome, String cidade, String endereco, String cep,
			Integer offset) {

		SelectConditionStep<Record4<String, String, String, String>> query = this.dslContext
				.select(PDVS.NOME, PDVS.CIDADE, PDVS.ENDERECO, PDVS.CEP).from(PDVS).where(PDVS.TENANT.eq(tenantId));

		if (StringUtils.trimToNull(nome) != null) {
			query.and(PDVS.NOME.eq(nome.trim()));
		}
		if (StringUtils.trimToNull(cidade) != null) {
			query.and(PDVS.CIDADE.eq(cidade.trim()));
		}
		if (StringUtils.trimToNull(endereco) != null) {
			query.and(PDVS.ENDERECO.eq(endereco.trim()));
		}
		if (StringUtils.trimToNull(cep) != null) {
			query.and(PDVS.CEP.eq(cep.trim()));
		}

		query.limit(1000);
		if (offset != null) {
			query.offset(offset);
		}

		return query.fetch(new RecordMapper<Record4<String, String, String, String>, PdvBean>() {

			@Override
			public PdvBean map(Record4<String, String, String, String> record) {
				PdvBean pdvBean = new PdvBean();

				pdvBean.nome = record.get(PDVS.NOME);
				pdvBean.cidade = record.get(PDVS.CIDADE);
				pdvBean.endereco = record.get(PDVS.ENDERECO);
				pdvBean.cep = record.get(PDVS.CEP);

				return pdvBean;
			}
		});
	}

	public static class PdvBean {
		private String nome;
		private String cidade;
		private String endereco;
		private String cep;

		public String getNome() {
			return nome;
		}

		public void setNome(String nome) {
			this.nome = nome;
		}

		public String getCidade() {
			return cidade;
		}

		public void setCidade(String cidade) {
			this.cidade = cidade;
		}

		public String getEndereco() {
			return endereco;
		}

		public void setEndereco(String endereco) {
			this.endereco = endereco;
		}

		public String getCep() {
			return cep;
		}

		public void setCep(String cep) {
			this.cep = cep;
		}

	}

}
