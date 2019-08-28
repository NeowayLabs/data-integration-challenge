package pdv.l3ug1m.github.com.pdvservice.controllers;

import java.nio.file.Path;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import pdv.l3ug1m.github.com.pdvservice.services.PdvFilesService;

@Controller
public class PdvsController {
	
	private static final org.slf4j.Logger log = org.slf4j.LoggerFactory.getLogger(PdvsController.class);

	@Autowired PdvFilesService pdvFilesService;
	
	
	@PostMapping("/uploadFile")
    public String handleFileUpload(@RequestParam("data") MultipartFile file,
    		@RequestParam("tenant") String tenantId, 
            RedirectAttributes redirectAttributes) {

		Path pathFile = pdvFilesService.store(file, tenantId);
		pdvFilesService.processFile(pathFile, tenantId);
        redirectAttributes.addFlashAttribute("message",
                "You successfully uploaded " + file.getOriginalFilename() + "!");

        return "redirect:/";
    }
	
	@ExceptionHandler
	public ResponseEntity<?> handleStorageFileNotFound(Throwable exc) {
		log.error(exc.getMessage(), exc);
		return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
	}
	
}
